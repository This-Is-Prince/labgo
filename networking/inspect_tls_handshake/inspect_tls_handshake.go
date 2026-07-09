package inspect_tls_handshake

import (
	"crypto/tls"
	"fmt"
)

func RunInspectTlsHandshake(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("Inspect TLS handshake\n\n")

	conn, err := tls.Dial("tcp", "api.eventsease.in:443", &tls.Config{
		ServerName: "api.eventsease.in",
		NextProtos: []string{"h2", "http/1.1"},
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()

	fmt.Println("Handshake complete:", state.HandshakeComplete)
	fmt.Println("Negotiated protocol:", state.NegotiatedProtocol)
	fmt.Println("Cipher suite:", tls.CipherSuiteName(state.CipherSuite))
	fmt.Println("Server name:", state.ServerName)
	fmt.Println("Version:", state.Version)

	for i, cert := range state.PeerCertificates {
		fmt.Println("Certificate", i)
		fmt.Println("  Subject:", cert.Subject)
		fmt.Println("  Issuer:", cert.Issuer)
		fmt.Println("  DNS names:", cert.DNSNames)
		fmt.Println("  Not before:", cert.NotBefore)
		fmt.Println("  Not after:", cert.NotAfter)
	}
}
