package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Practice bufio Package")

	f, err := os.Open("./index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := os.Remove("./README.md"); err != nil {
		fmt.Println(err)
	}
	readme, err := os.Create("./README.md")

	if err != nil {
		panic(err)
	}
	fmt.Println(readme)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		PrintH1(string(line))
	}
}
func PrintH1(line string) {
	_, after, found := strings.Cut(line, "<h1>")
	if found {
		before, _, _ := strings.Cut(after, "</h1>")
		fmt.Println(before)
	}
}
