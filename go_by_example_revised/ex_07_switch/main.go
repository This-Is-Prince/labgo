package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch")

	i := 1
	switch i {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	default:
		fmt.Println("Another")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Int")
		case float32:
			fmt.Println("Float32")
		case bool:
			fmt.Println("Bool")
		default:
			fmt.Printf("I don't know %T", t)
		}
	}
	whatAmI(false)
	whatAmI(int(2))
	whatAmI(float32(4.2))
	whatAmI("Hello")
}
