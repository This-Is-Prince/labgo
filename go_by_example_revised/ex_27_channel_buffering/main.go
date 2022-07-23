package main

import "fmt"

func main() {
	fmt.Println("Channel Buffering")

	msgs := make(chan string, 2)
	msgs <- "Hello World!"
	msgs <- "Namaste JavaScript"
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}
