package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Writing Files")

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("./tmp/dat.txt", d1, 0644)
	checkError(err)

	f, err := os.Create("./tmp/dat1.txt")
	checkError(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
