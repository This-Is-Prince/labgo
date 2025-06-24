package main

import "fmt"

func main() {
	fmt.Println("Anatomy of Conditional Statements and Loops in Go")

	condition := true

	if condition {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}

	fruit := "mangoa"

	if fruit == "apple" {
		fmt.Println("The fruit is an apple")
	} else if fruit == "banana" {
		fmt.Println("The fruit is a banana")
	} else if fruit == "mango" {
		fmt.Println("The fruit is a mango")
	} else {
		fmt.Println("Unknown fruit")
	}

	// This fruit scope is limited to the if statement
	if fruit := "apple"; fruit == "apple" {
		fmt.Println("The fruit is an apple")
	} else {
		fmt.Println("The fruit is not an apple")
	}

	// Switch statement
	switch fruit {
	case "apple":
		fmt.Println("The fruit is an apple")
	case "banana":
		fmt.Println("The fruit is a banana")
	case "mango":
		fmt.Println("The fruit is a mango")
	default:
		fmt.Println("Unknown fruit")
	}

	switch fruit {
	case "apple":
		{
			fmt.Println("The fruit is an apple")
			fmt.Println("Case block for apple")
		}
	case "banana":
		{
			fmt.Println("The fruit is a banana")
			fmt.Println("Case block for banana")
		}
	case "mango":
		{
			fmt.Println("The fruit is a mango")
			fmt.Println("Case block for mango")
		}
	default:
		{
			fmt.Println("Unknown fruit")
			fmt.Println("Default case block")
		}
	}

	switch fruit {
	default:
		{
			fmt.Println("Unknown fruit")
			fmt.Println("Default case block is in first position")
		}
	case "apple":
		{
			fmt.Println("The fruit is an apple")
			fmt.Println("Case block for apple")
		}
	case "banana":
		{
			fmt.Println("The fruit is a banana")
			fmt.Println("Case block for banana")
		}
	case "mango":
		{
			fmt.Println("The fruit is a mango")
			fmt.Println("Case block for mango")
		}
	}

	letter := "i"

	switch letter {
	case "a", "e", "i", "o", "u":
		{
			fmt.Println("The letter is a vowel")
		}
	default:
		{
			fmt.Println("The letter is a consonant")
		}
	}

	switch letter := "f"; letter {
	case "a", "e", "i", "o", "u":
		{
			fmt.Println("The letter is a vowel")
		}
	default:
		{
			fmt.Println("The letter is a consonant")
		}
	}

	switch {
	case letter == "a":
		{
			fmt.Println("The letter is a")
		}
	default:
		{
			fmt.Println("The letter is not a")
		}
	}

	switch number := getNumber(); {
	case number < 10:
		fmt.Println("The number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("The number is between 10 and 20")
		fallthrough
	case number >= 20 && number < 30:
		fmt.Println("The number is between 20 and 30")
		fallthrough
	case number >= 30 && number < 40:
		fmt.Println("The number is between 30 and 40")
	default:
		fmt.Println("The number is 40 or greater")
	}

	for i := 1; i <= 6; i++ {
		fmt.Println("Iteration number:", i)
	}

	for i := 1; i <= 6; {
		fmt.Println("Iteration number:", i)
		i++
	}

	j := 1
	for j < 6 {
		fmt.Println("Iteration number:", j)
		j++
	}

	i := 1
	for {
		fmt.Println("Iteration number:", i)
		if i >= 5 {
			break
		}
		i++
	}

	fmt.Println("Program Terminated... exiting main function")
}

func getNumber() int {
	return 32
}
