package tcpclient

import (
	"fmt"
	"net"
)

func RunTcpClient(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("TCP client\n\n")

	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	request := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"

	_, err = conn.Write([]byte(request))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err != nil {
			break
		}
	}
}
