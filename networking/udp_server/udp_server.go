package udpserver

import (
	"fmt"
	"net"
)

func RunUdpServer(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("UDP server\n\n")

	addr, err := net.ResolveUDPAddr("udp", ":9000")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on :9000")

	buf := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		fmt.Printf("from %s: %s\n", clientAddr, string(buf[:n]))

		conn.WriteToUDP([]byte("received\n"), clientAddr)
	}
}
