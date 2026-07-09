package h2frame_client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

func RunH2FrameClient(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("H2 Frame Client\n\n")

	rawConn, err := net.Dial("tcp", "127.0.0.1:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer rawConn.Close()

	tlsConn := tls.Client(rawConn, &tls.Config{
		ServerName:         "localhost",
		InsecureSkipVerify: true, // local self-signed lab only
		NextProtos:         []string{"h2"},
		KeyLogWriter:       keyLogWriter(),
	})
	defer tlsConn.Close()

	if err := tlsConn.Handshake(); err != nil {
		log.Fatal(err)
	}

	state := tlsConn.ConnectionState()

	fmt.Println("TLS negotiated protocol:", state.NegotiatedProtocol)
	fmt.Println("Cipher suite:", tls.CipherSuiteName(state.CipherSuite))

	if state.NegotiatedProtocol != "h2" {
		log.Fatal("HTTP/2 was not negotiated")
	}

	fr := http2.NewFramer(tlsConn, tlsConn)

	fmt.Println()
	fmt.Println("WRITE: HTTP/2 client connection preface")
	fmt.Printf("%q\n", http2.ClientPreface)

	_, err = tlsConn.Write([]byte(http2.ClientPreface))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("WRITE FRAME: SETTINGS")
	if err := fr.WriteSettings(); err != nil {
		log.Fatal(err)
	}

	// Read server SETTINGS and ACK them.
	for {
		frame, err := fr.ReadFrame()
		if err != nil {
			log.Fatal(err)
		}

		printFrame("READ FRAME", frame)

		if sf, ok := frame.(*http2.SettingsFrame); ok && !sf.IsAck() {
			fmt.Println("WRITE FRAME: SETTINGS ACK")
			if err := fr.WriteSettingsAck(); err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	var headerBlock bytes.Buffer
	encoder := hpack.NewEncoder(&headerBlock)

	headers := []hpack.HeaderField{
		{Name: ":method", Value: "POST"},
		{Name: ":scheme", Value: "https"},
		{Name: ":authority", Value: "localhost:8443"},
		{Name: ":path", Value: "/v1/events?source=h2lab"},
		{Name: "content-type", Value: "application/json"},
		{Name: "accept", Value: "application/json"},
	}

	fmt.Println()
	fmt.Println("HTTP/2 request headers before HPACK encoding:")
	for _, h := range headers {
		fmt.Printf("  %s: %s\n", h.Name, h.Value)

		if err := encoder.WriteField(h); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println()
	fmt.Println("WRITE FRAME: HEADERS stream=1")
	if err := fr.WriteHeaders(http2.HeadersFrameParam{
		StreamID:      1,
		BlockFragment: headerBlock.Bytes(),
		EndHeaders:    true,
		EndStream:     false,
	}); err != nil {
		log.Fatal(err)
	}

	body := []byte(`{"title":"Music Night"}`)

	fmt.Println("WRITE FRAME: DATA stream=1 endStream=true")
	fmt.Println("DATA payload:", string(body))

	if err := fr.WriteData(1, true, body); err != nil {
		log.Fatal(err)
	}

	decoder := hpack.NewDecoder(4096, func(f hpack.HeaderField) {
		fmt.Printf("  %s: %s\n", f.Name, f.Value)
	})

	for {
		frame, err := fr.ReadFrame()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}

		printFrame("READ FRAME", frame)

		switch f := frame.(type) {
		case *http2.HeadersFrame:
			fmt.Println("Decoded response headers:")
			if _, err := decoder.Write(f.HeaderBlockFragment()); err != nil {
				log.Println("HPACK decode error:", err)
			}

			if f.StreamEnded() {
				return
			}

		case *http2.DataFrame:
			fmt.Println("Response DATA payload:")
			fmt.Println(string(f.Data()))

			if f.StreamEnded() {
				return
			}
		}
	}
}

func printFrame(prefix string, frame http2.Frame) {
	h := frame.Header()

	fmt.Printf(
		"%s: type=%v stream=%d flags=%v length=%d\n",
		prefix,
		h.Type,
		h.StreamID,
		h.Flags,
		h.Length,
	)
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
