package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("---------Race Condition---------")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)
	wg.Wait()
	mut.RLock()
	fmt.Println("Score is:-", score)
	mut.RUnlock()
}
