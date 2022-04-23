package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var one int = 2
	fmt.Println("Before Changing One:-", one)
	changeOne(one)
	fmt.Println("After Changing One:-", one)

	var onePtr *int = &one
	fmt.Println("Before Changing onePtr:-", *onePtr)
	changeOnePtr(onePtr)
	fmt.Println("After Changing onePtr:-", *onePtr)

	/* 	var ptr *int
	   	fmt.Println("Value of pointer is ", ptr) */

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}

func changeOne(one int) {
	one++
}
func changeOnePtr(onePtr *int) {
	*onePtr++
}
