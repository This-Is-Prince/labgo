package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Working with files and file system: A low-level introduction")

	// Os_Create()
	// Os_Open()
	// Os_OpenFile()
	// Os_Symlink()
	WorkingWithFiles()
}

func Os_Create() {
	// file, err := os.Create("./tmp/temp.txt") // ERROR
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func Os_Open() {
	file, err := os.Open("./temp1.txt")
	if err != nil {
		fmt.Println(os.IsExist(err))
		fmt.Println(os.IsNotExist(err))
	}
	defer file.Close()
}

func Os_OpenFile() {
	file, err := os.OpenFile("./temp.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("Hello World!")
	file.ReadFrom(os.Stdin)
}

func Os_Symlink() {
	err := os.Symlink("./temp.txt", "./tmp")
	if err != nil {
		log.Fatal(err)
	}
}

func WorkingWithFiles() {
	// open a `info.txt` file in current working directory
	// create the file if doesn't exist and open it for reading and writing
	// append the content to the file while writing instead of overriding
	// create the file in 0744 file-mode (read-only for others)
	file, err := os.OpenFile("info.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)

	// log error if exist
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file.Close()

	// append new content
	fmt.Fprint(file, "Hello World!\n")
	fmt.Fprintln(file, "How are you?")
	fmt.Fprintf(file, "%s am %s\n", "I", "Prince")
}
