package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One Routine")

		m.Lock()
		score = append(score, 1)
		m.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two Routine")

		m.Lock()
		score = append(score, 2)
		m.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three Routine")

		m.Lock()
		score = append(score, 3)
		m.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four Routine")

		m.RLock()
		fmt.Println(score)
		m.RUnlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()

	mut.RLock()
	fmt.Println(score)
	mut.RUnlock()
}
