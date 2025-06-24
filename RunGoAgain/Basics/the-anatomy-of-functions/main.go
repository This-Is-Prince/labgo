package main

import "fmt"

func main() {
	fmt.Println("The anatomy of Functions in Go!")

	doSomething()
	greet("John Doe")
	add(1, 2)
	addFloat(2.2, 1.2)
	add, mul := addMul(2.2, 1.2)
	fmt.Println("Add, Mul", add, mul)
	add, mul = nameReturn(2.2, 1.2)
	fmt.Println("Add, Mul", add, mul)

	defer sayOne()
	fmt.Println("Main Finished")

	fmt.Printf("Type of add function %T\n", add1)
	fmt.Printf("Type of subtract function %T\n", subtract1)

	calc(1, 2, add1)
	calc(1, 2, subtract1)

	calc1(1, 2, add1)
	calc1(1, 2, subtract1)

	add2(1.2, 2.1)
	func(a, b string) {
		fmt.Println("Add of 2 string", a+b)
	}("Hi", "Ji")
}

func doSomething() {
	fmt.Println("Hello World!")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func add(a, b int) {
	fmt.Println("Sum is:- ", a+b)
}

func addFloat(a, b float32) float32 {
	add := a + b
	fmt.Println("Add Float is", add)
	return add
}

func addMul(a, b float32) (float32, float32) {
	return a + b, a * b
}

func nameReturn(a, b float32) (add float32, mul float32) {
	add = a + b
	mul = a * b

	return
}

func sayOne() {
	fmt.Println("I am One")
}

func add1(a, b int) int {
	fmt.Println("Addition is:-", a+b)
	return a + b
}

func subtract1(a, b int) int {
	fmt.Println("Subtraction is:-", a-b)
	return a - b
}

func calc(a, b int, f func(int, int) int) {
	fmt.Println("output of calc is", f(a, b))
}

// Derived Type
type CalcFunc func(int, int) int

func calc1(a, b int, f CalcFunc) {
	fmt.Println("Output of calc1 is", f(a, b))
}

var add2 = func(a, b float32) float32 {
	return a + b
}
