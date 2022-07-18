package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(fact(5))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
