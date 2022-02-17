package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// world, One, Two

func myDefer() {
	for i := 0; i < 5; i++ {
		if i == 0 {
			defer fmt.Println(i)
		} else {
			defer fmt.Print(i)
		}
	}
}
