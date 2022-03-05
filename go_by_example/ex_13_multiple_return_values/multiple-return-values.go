package main

import "fmt"

func returnMultipleValues() (int, int) {
	return 3, 7
}

func main() {
	fmt.Println("---------Multiple Return Values---------")

	a, b := returnMultipleValues()
	fmt.Println(a)
	fmt.Println(b)

	_, c := returnMultipleValues()
	fmt.Println(c)
}
