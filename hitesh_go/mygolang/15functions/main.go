package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println("Pro Result is: ", proRes)

	// function declaration is not defined inside another function
	// func greeterTwo(){
	// 	fmt.Println("Another method")
	// }
	// greeterTwo()
}

// func greeterTwo() {
// 	fmt.Println("Another method")
// }

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namstey from golang")
}
