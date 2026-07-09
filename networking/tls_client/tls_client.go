package tlsclient

import (
	"crypto/tls"
	"fmt"
)

func RunTlsClient(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("Running TLS client\n\n")

	conn, err := tls.Dial("tcp", "api.eventsease.in:443", &tls.Config{
		ServerName: "api.eventsease.in",
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()

	fmt.Println("TLS version:", state.Version)
	fmt.Println("Negotiated protocol:", state.NegotiatedProtocol)

	fmt.Println()

	for _, cert := range state.PeerCertificates {
		fmt.Println("certificate subject:", cert.Subject)
		fmt.Println("issuer:", cert.Issuer)
		fmt.Println("dns names:", cert.DNSNames)

		fmt.Println()
	}
}
