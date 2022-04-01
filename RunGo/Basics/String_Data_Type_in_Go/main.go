package main

import "fmt"

func main() {
	fmt.Println("String Data Type in Go")

	simpleProgram()
	lengthOfString()
	asciiCharacter()
}

func simpleProgram() {
	var s string

	s = "Hello World"

	fmt.Println(s)
}

func lengthOfString() {
	s := "Hello World"
	fmt.Println(len(s))
}

func asciiCharacter() {
	s := "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}
	fmt.Println()
}
