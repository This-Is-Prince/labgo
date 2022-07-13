package main

import "fmt"

func main() {
	i := 0

	fmt.Print("First: ")
	for i <= 10 {
		fmt.Print(i, ",")
		i++
	}
	fmt.Print("\nSecond: ")
	for i := 0; i <= 10; i++ {
		fmt.Print(i, ",")
	}
	fmt.Println("\nThird: ")
	j := 0
	for {
		if j == 10 {
			break
		}
		fmt.Print(i, ",")
		j++
	}
}
