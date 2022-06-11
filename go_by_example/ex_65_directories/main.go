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
	fmt.Println("Directories")

	err := os.Mkdir("subdir", 0755)
	checkError(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		checkError(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1.txt")

	err = os.MkdirAll("subdir/parent/child", 0755)
	checkError(err)

	createEmptyFile("subdir/parent/file2.txt")
	createEmptyFile("subdir/parent/file3.txt")
	createEmptyFile("subdir/parent/child/file4.txt")

	c, err := os.ReadDir("subdir/parent")
	checkError(err)

	fmt.Println("Listing subdir/parent")

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	checkError(err)

	c, err = os.ReadDir(".")
	checkError(err)
	fmt.Println("Listing subdir/parent/child")

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	checkError(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	checkError(err)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
