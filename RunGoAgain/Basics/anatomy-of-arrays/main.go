package main

import "fmt"

func main() {
	fmt.Println("The anatomy of Arrays in Go!")

	var a [3]int          // Declare an array of 3 integers
	a[0] = 1              // Assign value to the first element
	fmt.Println("a =", a) // Print the first element

	var b [3]int = [3]int{1, 2, 3}
	fmt.Println(b)

	var c = [3]int{1, 2, 3}
	fmt.Println(c)

	d := [3]int{1, 2, 3}
	fmt.Println(d)

	e := [3]int{1, 2}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 45}
	fmt.Println(f, len(f))

	greetings := [3]string{
		"Hello!",
		"Hi!",
		"Hey!",
	}
	fmt.Println(greetings)

	a = [3]int{1, 2, 3}
	b = [3]int{1, 1, 1}
	c = [3]int{1, 1, 2}
	d = [3]int{1, 2, 3}
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(a == d)

	for range 10 {
		fmt.Println("Hello, World!")
	}

	for i, v := range a {
		fmt.Println(i, v)
	}

	m := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(m)
}
