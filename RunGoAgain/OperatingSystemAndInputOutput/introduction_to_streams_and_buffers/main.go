package main

import (
	"fmt"
	"io"
	"strings"
)

type MyStringData struct {
	str       string
	readIndex int
}

func (myStringData *MyStringData) Read(p []byte) (n int, err error) {
	strBytes := []byte(myStringData.str)

	if myStringData.readIndex >= len(strBytes) {
		return 0, io.EOF
	}

	nextReadLimit := myStringData.readIndex + len(p)

	if nextReadLimit >= len(strBytes) {
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	nextBytes := strBytes[myStringData.readIndex:nextReadLimit]
	n = len(nextBytes)

	copy(p, nextBytes)
	myStringData.readIndex = nextReadLimit

	return
}

func main() {
	src := MyStringData{
		str: "Hello Amazing World!",
	}

	p := make([]byte, 3)

	for {
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}

	src1 := strings.NewReader("Hello Amazing World!")

	p1 := make([]byte, 3)

	for {
		n, err := src1.Read(p1)
		fmt.Printf("%d bytes read, data: %s\n", n, p1[:n])

		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}

	src2 := strings.NewReader("Hello Amazing World!")
	data, err := io.ReadAll(src2)
	fmt.Printf("Read data of length %d : %s\n", len(data), data)
	fmt.Println(err)

	src3 := strings.NewReader("Hello Amazing World!")
	buf := make([]byte, 14)

	bytesRead1, err1 := io.ReadFull(src3, buf)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead1, buf[:bytesRead1], err1)

	bytesRead2, err2 := io.ReadFull(src3, buf)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead2, buf[:bytesRead2], err2)

	bytesRead3, err3 := io.ReadFull(src3, buf)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead3, buf[:bytesRead3], err3)
}
