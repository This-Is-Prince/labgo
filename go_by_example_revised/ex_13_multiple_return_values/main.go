package main

import "fmt"

func main() {
	fmt.Println("Multiple Return Values")
	fmt.Println(values())
}

func values() (int, int) {
	return 3, 7
}
