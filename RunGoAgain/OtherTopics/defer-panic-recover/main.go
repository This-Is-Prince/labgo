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

func d() {
	fmt.Println("Function d is called.")

	defer func() {
		fmt.Println("Deferred function in d is called.")
		defer e()
		fmt.Println("Deferred function in d is returning.")
	}()
	panic("A panic occurred in function d!")

	fmt.Println("Function d is returning.")
}

func e() {
	fmt.Println("Function e is called.")

	if r := recover(); r != nil {
		fmt.Println("Recovered in function e:", r)
	} else {
		fmt.Println("No panic occurred in function e.")
	}

	fmt.Println("Function e is returning.")
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
	d()
	c()
	b()

	panic("A panic has occurred!")

	fmt.Println("Ending the program...")
}
