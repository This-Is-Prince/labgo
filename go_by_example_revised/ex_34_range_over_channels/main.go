package main

import "fmt"

func main() {
	fmt.Println("Range Over Channels")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
