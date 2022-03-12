package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("------- Channels and Deadlock in golang -------")
	/*
		myCh := make(chan int)
		wg := &sync.WaitGroup{}

		// myCh <- 5 // ERROR, all goroutines are asleep Deadlock
		// fmt.Println(<-myCh)

		wg.Add(2)
		go func(ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(<-ch)
		}(myCh, wg)

		go func(ch chan int, wg *sync.WaitGroup) {
			ch <- 5
			ch <- 6 // ERROR, all goroutines are asleep Deadlock
			defer wg.Done()
		}(myCh, wg)
		wg.Wait() */

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, isChannelOpen := <-ch
		fmt.Println(val, isChannelOpen, len(ch))
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 5
		ch <- 6
		close(ch)
	}(myCh, wg)
	wg.Wait()
}
