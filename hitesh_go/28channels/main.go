package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")
	/*
		myCh := make(chan int)
		wg := &sync.WaitGroup{}

		// fmt.Println(<-myCh)
		// myCh <- 5

		wg.Add(2)
		go func(ch chan int, wg *sync.WaitGroup) {
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			wg.Done()
		}(myCh, wg)

		go func(ch chan int, wg *sync.WaitGroup) {
			ch <- 5
			ch <- 6
			wg.Done()
		}(myCh, wg)

		wg.Wait() */

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 0
		// ch <- 6
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
