package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
	fmt.Println("Program started at", start)
}

func PrintHello() {
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Hello World!")
}

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits(s []int) {
	for _, d := range s {
		fmt.Printf("%d at time %v\n", d, time.Since(start))
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Anatomy of goroutines in Go - Concurrency in Go")

	// printHello()
	// go PrintHello() // This will run in a separate goroutine

	// time.Sleep(10 * time.Millisecond)
	go func() {
		fmt.Println("Hey")
	}()

	go getChars("Hello")

	go getDigits([]int{1, 2, 3, 4, 5})

	time.Sleep(100 * time.Millisecond)

	fmt.Println("main execution stopped at time", time.Since(start))
}
