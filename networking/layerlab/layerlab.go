package layerlab

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func RunLayerLab(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("Layer Lab\n\n")

	mode := flag.String("mode", "", "plain-server | plain-client | tls-server | tls-client")
	flag.Parse()

	switch *mode {
	case "plain-server":
		runPlainServer()
	case "plain-client":
		runPlainClient()
	case "tls-server":
		runTLSServer()
	case "tls-client":
		runTLSClient()
	default:
		fmt.Println("usage:")
		fmt.Println("  go run layerlab.go -mode plain-server")
		fmt.Println("  go run layerlab.go -mode plain-client")
		fmt.Println("  go run layerlab.go -mode tls-server")
		fmt.Println("  go run layerlab.go -mode tls-client")
	}
}

// -----------------------------
// Pretty byte dumping
// -----------------------------

func dumpBytes(title string, b []byte) {
	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println(title)
	fmt.Println("bytes:", len(b))
	fmt.Println("==================================================")

	if len(b) > 1024 {
		fmt.Println(hex.Dump(b[:1024]))
		fmt.Printf("... truncated, total %d bytes\n", len(b))
		return
	}

	fmt.Println(hex.Dump(b))
}

// -----------------------------
// Lab 1: plain TCP server/client
// -----------------------------

func runPlainServer() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("plain TCP server listening on 127.0.0.1:8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			fmt.Println("accepted connection from:", c.RemoteAddr())

			buf := make([]byte, 8192)

			n, err := c.Read(buf)
			if err != nil {
				log.Println("read error:", err)
				return
			}

			dumpBytes("SERVER RECEIVED RAW TCP BYTES: plaintext HTTP request", buf[:n])

			response := "" +
				"HTTP/1.1 200 OK\r\n" +
				"Content-Type: application/json\r\n" +
				"Content-Length: 15\r\n" +
				"Connection: close\r\n" +
				"\r\n" +
				`{"ok":true}` + "\n"

			dumpBytes("SERVER WRITES RAW TCP BYTES: plaintext HTTP response", []byte(response))

			_, _ = c.Write([]byte(response))
		}(conn)
	}
}

func runPlainClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	body := `{"title":"Music Night"}`

	request := fmt.Sprintf(
		"POST /v1/events?source=lab HTTP/1.1\r\n"+
			"Host: localhost:8080\r\n"+
			"Content-Type: application/json\r\n"+
			"Accept: application/json\r\n"+
			"Content-Length: %d\r\n"+
			"Connection: close\r\n"+
			"\r\n"+
			"%s",
		len(body),
		body,
	)

	dumpBytes("CLIENT WRITES RAW TCP BYTES: plaintext HTTP request", []byte(request))

	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	dumpBytes("CLIENT RECEIVED RAW TCP BYTES: plaintext HTTP response", resp)
}

// -----------------------------
// Lab 2: TLS encrypted transport
// -----------------------------

type loggingConn struct {
	net.Conn
	name string
}

func (c *loggingConn) Read(p []byte) (int, error) {
	n, err := c.Conn.Read(p)
	if n > 0 {
		dumpBytes(c.name+" TCP READ: encrypted bytes from socket", p[:n])
	}
	return n, err
}

func (c *loggingConn) Write(p []byte) (int, error) {
	if len(p) > 0 {
		dumpBytes(c.name+" TCP WRITE: encrypted bytes to socket", p)
	}
	return c.Conn.Write(p)
}

type loggingListener struct {
	net.Listener
	name string
}

func (l *loggingListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	fmt.Println("accepted TCP connection from:", conn.RemoteAddr())

	return &loggingConn{
		Conn: conn,
		name: l.name,
	}, nil
}

func runTLSServer() {
	cert, err := generateSelfSignedCert()
	if err != nil {
		log.Fatal(err)
	}

	baseLn, err := net.Listen("tcp", "127.0.0.1:8443")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TLS server listening on https://127.0.0.1:8443")

	ln := &loggingListener{
		Listener: baseLn,
		name:     "SERVER-UNDERLYING",
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},

		// Try HTTP/2 first, then HTTP/1.1.
		// Our manual client below forces HTTP/1.1 for easier raw viewing.
		NextProtos: []string{"h2", "http/1.1"},

		// If SSLKEYLOGFILE is set, Go writes TLS secrets there.
		// Wireshark can use that file to decrypt TLS.
		KeyLogWriter: keyLogWriter(),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		fmt.Println("==================================================")
		fmt.Println("SERVER HANDLER: decrypted HTTP request")
		fmt.Println("==================================================")
		fmt.Println("Proto:", r.Proto)
		fmt.Println("Method:", r.Method)
		fmt.Println("Host:", r.Host)
		fmt.Println("Path:", r.URL.Path)
		fmt.Println("RawQuery:", r.URL.RawQuery)

		fmt.Println("Headers:")
		for k, v := range r.Header {
			fmt.Printf("  %s: %v\n", k, v)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}

		dumpBytes("SERVER HANDLER: decrypted HTTP body", body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"evt_123","ok":true}`))
	})

	server := &http.Server{
		Handler: handler,
	}

	err = server.Serve(tls.NewListener(ln, tlsConfig))
	if err != nil {
		log.Fatal(err)
	}
}

func runTLSClient() {
	rawConn, err := net.Dial("tcp", "127.0.0.1:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer rawConn.Close()

	loggedRawConn := &loggingConn{
		Conn: rawConn,
		name: "CLIENT-UNDERLYING",
	}

	tlsConn := tls.Client(loggedRawConn, &tls.Config{
		ServerName: "localhost",

		// For local self-signed lab only.
		// Real browser verification uses trusted CA roots.
		InsecureSkipVerify: true,

		// Force HTTP/1.1 so we can manually write readable HTTP text.
		NextProtos: []string{"http/1.1"},

		KeyLogWriter: keyLogWriter(),
	})
	defer tlsConn.Close()

	fmt.Println("starting TLS handshake...")

	if err := tlsConn.Handshake(); err != nil {
		log.Fatal(err)
	}

	state := tlsConn.ConnectionState()

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println("CLIENT TLS CONNECTION STATE")
	fmt.Println("==================================================")
	fmt.Println("Handshake complete:", state.HandshakeComplete)
	fmt.Println("Negotiated protocol:", state.NegotiatedProtocol)
	fmt.Println("Cipher suite:", tls.CipherSuiteName(state.CipherSuite))
	fmt.Println("TLS version:", tlsVersionName(state.Version))

	for i, cert := range state.PeerCertificates {
		fmt.Println("Certificate", i)
		fmt.Println("  Subject:", cert.Subject)
		fmt.Println("  Issuer:", cert.Issuer)
		fmt.Println("  DNSNames:", cert.DNSNames)
		fmt.Println("  NotBefore:", cert.NotBefore)
		fmt.Println("  NotAfter:", cert.NotAfter)
	}

	body := `{"title":"Music Night"}`

	request := fmt.Sprintf(
		"POST /v1/events?source=lab HTTP/1.1\r\n"+
			"Host: localhost:8443\r\n"+
			"Content-Type: application/json\r\n"+
			"Accept: application/json\r\n"+
			"Content-Length: %d\r\n"+
			"Connection: close\r\n"+
			"\r\n"+
			"%s",
		len(body),
		body,
	)

	dumpBytes("CLIENT: plaintext HTTP request JUST BEFORE TLS ENCRYPTION", []byte(request))

	_, err = tlsConn.Write([]byte(request))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := io.ReadAll(tlsConn)
	if err != nil {
		log.Fatal(err)
	}

	dumpBytes("CLIENT: plaintext HTTP response AFTER TLS DECRYPTION", resp)
}

func generateSelfSignedCert() (tls.Certificate, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore: time.Now().Add(-time.Hour),
		NotAfter:  time.Now().Add(24 * time.Hour),

		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},

		DNSNames: []string{"localhost"},
		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
		},
	}

	certDER, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&privateKey.PublicKey,
		privateKey,
	)
	if err != nil {
		return tls.Certificate{}, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return tls.X509KeyPair(certPEM, keyPEM)
}

func keyLogWriter() io.Writer {
	path := os.Getenv("SSLKEYLOGFILE")
	if path == "" {
		return nil
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Println("could not open SSLKEYLOGFILE:", err)
		return nil
	}

	fmt.Println("writing TLS key log to:", path)
	return f
}

func tlsVersionName(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("unknown: %d", v)
	}
}
