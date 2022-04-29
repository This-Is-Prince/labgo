package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Prince"

	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length == len(content))

	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databyte)
	fmt.Println(string(databyte))
}
