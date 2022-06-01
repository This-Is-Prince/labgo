package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Concurrency using go routines and Channels.")
	started := time.Now()
	foods := []string{"mashroom pizza", "pasta", "kebab", "cake"}
	str := ""

	/*
	   results := make(chan int)
	   	for _, f := range foods {
	   		go func(f string) {
	   			cook(f)
	   			value, isPresent := <-results
	   			if isPresent {
	   				str += strconv.FormatInt(int64(value), 32) + ", "
	   			}
	   		}(f)
	   	}
	   	for i := 0; i < len(foods); i++ {
	   		results <- i
	   	}
	*/

	results := make(chan int, 3)
	for _, f := range foods {
		go func(f string) {
			cook(f)
			value, isPresent := <-results
			if isPresent {
				str += strconv.FormatInt(int64(value), 32) + ", "
			}
		}(f)
	}
	for i := 0; i < len(foods); i++ {
		results <- i
	}
	fmt.Printf("done in %v\n", time.Since(started))
	fmt.Println(str)
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
}
