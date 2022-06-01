package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Concurrency & WaitGroup")

	started := time.Now()
	foods := []string{"mashroom pizza", "pasta", "kebab", "cake"}

	var wg sync.WaitGroup
	wg.Add(len(foods))
	for _, f := range foods {
		go func(f string) {
			cook(f)
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Printf("done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}
