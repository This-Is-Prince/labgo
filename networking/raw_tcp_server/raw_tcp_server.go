package rawtcpserver

import (
	"fmt"
	"io"
	"net"
)

func RunRawTcpServer(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("Raw TCP Server\n\n")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	fmt.Println("new connection from:", conn.RemoteAddr())

	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			return
		}

		fmt.Println("received:")
		fmt.Println(string(buf[:n]))

		conn.Write([]byte("hello from raw TCP server\n"))
	}
}
