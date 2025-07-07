package main

import "fmt"

func a() {
	fmt.Println("Function a is called.")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in function a:", r)
		} else {
			fmt.Println("No panic occurred in function a.")
		}
	}()

	panic("A panic occurred in function a!")

	fmt.Println("Function a is returning.")
}

func b() {
	fmt.Println("Function b is called.")

	panic("A panic occurred in function b!")

	fmt.Println("Function b is returning.")
}

func c() {
	fmt.Println("Function c is called.")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in function c:", r)
			if v, ok := r.(struct {
				Name string
				Age  int
			}); ok {
				fmt.Printf("Recovered struct: Name=%s, Age=%d\n", v.Name, v.Age)
			} else {
				fmt.Println("Recovered value is not of expected type.")
			}
		} else {
			fmt.Println("No panic occurred in function c.")
		}
	}()

	panic(struct {
		Name string
		Age  int
	}{
		Name: "John Doe",
		Age:  30,
	})

	fmt.Println("Function c is returning without panic.")
}

func main() {
	fmt.Println("Starting the program...")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		} else {
			fmt.Println("No panic occurred.")
		}
	}()

	a()
	c()
	b()

	panic("A panic has occurred!")

	fmt.Println("Ending the program...")
}
