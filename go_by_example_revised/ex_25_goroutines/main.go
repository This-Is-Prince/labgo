package main

import (
	"fmt"
	"time"
)

func printStringToCharacters(s string) {
	for _, v := range s {
		fmt.Print(string(v), ":", v, ", ")
	}
	fmt.Println()
}

func main() {
	printStringToCharacters("Hello World")
	go printStringToCharacters("Hello World")

	go func(msg string) {
		fmt.Println(msg)
	}("Hello World")

	time.Sleep(time.Second)
	fmt.Println("done")
}
