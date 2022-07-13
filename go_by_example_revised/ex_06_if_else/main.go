package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	for i := 0; i <= 10; i++ {
		displayIsEvenOrOdd(checkOddOrEven(i))
	}
}

func displayIsEvenOrOdd(isEven bool) {
	if isEven {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func checkOddOrEven(num int) bool {
	return num%2 == 0
}
