package main

import "fmt"

func main() {
	fmt.Println("Variadic Functions in Go!")

	s := []int{1, 2, 3, 4, 5}
	multiplies := getMultiplies(2, s...)
	fmt.Println(s, multiplies)
}

func getMultiplies(factor int, args ...int) []int {

	for i, v := range args {
		args[i] = factor * v
	}

	return args
}
