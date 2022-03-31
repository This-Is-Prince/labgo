package main

import "fmt"

func main() {
	fmt.Println("Anatomy of Conditional Statements and Loops in Go")

	theIfCondition()
	theIfElseCondition()
	theIfElseElseIfCondition()
	initialStatement()
	theSwitchConditionalStatement()
	expressionlessSwitchStatement()
	theForLoops()
	variantsOfForLoop()

}

func theIfCondition() {
	fmt.Println("The If Condition.")
	condition := true
	if condition {
		fmt.Println("condition met")
	}
	fmt.Println("program terminated")
}

func theIfElseCondition() {
	fmt.Println("The If/Else Condition")

	a := 2

	if a > 10 {
		fmt.Println("condition met")
	} else {
		fmt.Println("condition did not meet")
	}
	fmt.Println("program terminated")
}

func theIfElseElseIfCondition() {
	fmt.Println("The If/Else Else/If Condition")

	fruit := "orange"
	if fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}
}

func initialStatement() {
	fmt.Println("Initial Statement")

	if fruit := "banana"; fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}
}

func theSwitchConditionalStatement() {
	fmt.Println("The Switch Conditional Statement")

	finger := 2
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}

	// The Default Case
	fmt.Println("The Default Case")

	finger = 6
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("No fingers matched")
	}

	// Multiple case values
	fmt.Println("Multiple case values")
	letter := "i"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Letter is a Vowel.")
	default:
		fmt.Println("Letter is not a Vowel.")
	}

	// The initial statement
	fmt.Println("The initial Statement")

	switch letter := "b"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Letter is a Vowel.")
	default:
		fmt.Println("Letter is not a Vowel.")
	}
}

func expressionlessSwitchStatement() {
	fmt.Println("Expressionless Switch Statement")

	number := 20

	switch {
	case number <= 5:
		fmt.Println("number is less than or equal to 5.")
	case number > 5:
		fmt.Println("number is greater than 5.")
	case number > 10:
		fmt.Println("number is greater than 10.")
	case number > 15:
		fmt.Println("number greater than 15.")
	}

	switch number := getnumber(); {
	case number <= 5:
		fmt.Println("number is less than or equal to 5.")
	case number > 5:
		fmt.Println("number is greater than 5.")
	case number > 10:
		fmt.Println("number is greater than 10.")
	case number > 15:
		fmt.Println("number greater than 15.")
	}

	fmt.Println("The fallthrough statement")

	switch number := 20; {
	case number <= 5:
		fmt.Println("number is less than or equal to 5.")
		fallthrough
	case number > 5:
		fmt.Println("number is greater than 5.")
		fallthrough
	case number > 10:
		fmt.Println("number is greater than 10.")
		fallthrough
	case number > 15:
		fmt.Println("number greater than 15.")
	}
}

func getnumber() int {
	return 20
}

func theForLoops() {
	fmt.Println("The for Loops")

	for i := 1; i <= 6; i++ {
		fmt.Printf("Current number is %d \n", i)
	}

}

func variantsOfForLoop() {
	fmt.Println("Variants of the for Loop")

	fmt.Println("Optional post statement")
	for i := 1; i <= 6; {
		fmt.Printf("Current number is %d \n", i)
		i++
	}

	fmt.Println("Optional init statement")
	i := 1
	for ; i <= 6; i++ {
		fmt.Printf("Current number is %d \n", i)
	}

	fmt.Println("Optional init statement and post statement")
	i = 1
	for i <= 6 {
		fmt.Printf("Current number is %d \n", i)
		i++
	}

	fmt.Println("Without all statements.")
	i = 1
	for {
		fmt.Printf("Current number is %d \n", i)

		if i == 6 {
			break
		}
		i++
	}
}
