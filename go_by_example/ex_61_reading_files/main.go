package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fmt.Println("Reading Files")

	dat, err := os.ReadFile("./tmp/dat.txt")
	checkError(err)
	fmt.Println(string(dat))

	f, err := os.Open("./tmp/dat.txt")
	checkError(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(5, 1)
	checkError(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	checkError(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	checkError(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkError(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	checkError(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	checkError(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
