package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, value bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- value
}

func main() {
	fmt.Println("Channel Synchronization")

	done := make(chan bool, 1)

	for i := 0; i < 10; i++ {
		if i%7 == 0 {
			go worker(done, true)
		} else {
			go worker(done, false)
		}
	}

	for {
		if value := <-done; value {
			fmt.Println("This is true")
			break
		} else {
			fmt.Println("This is false")
		}
	}
}
