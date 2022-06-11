package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Temporary Files And Directories")

	f, err := os.CreateTemp("", "sample.txt")
	checkError(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())
	_, err = f.Write([]byte{1, 2, 3, 4})
	checkError(err)
	dname, err := os.MkdirTemp("", "sampledir")
	checkError(err)
	fmt.Println("Temp dir name:", dname)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1.txt")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	checkError(err)
}
