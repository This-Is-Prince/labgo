package main

import "fmt"

func main() {
	fmt.Println("Channels")

	msg := make(chan string)

	go func() {
		msg <- "Hello World!"
	}()

	value, isPresent := <-msg
	fmt.Println(value, isPresent)
}
