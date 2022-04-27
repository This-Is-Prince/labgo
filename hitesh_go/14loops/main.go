package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Print(days[d], ", ")
	}
	fmt.Println()

	for i := range days {
		fmt.Print(days[i], ", ")
	}
	fmt.Println()

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto prince
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		if rougueValue == 8 {
			break
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

prince:
	fmt.Println("Jumping at prince")
}
