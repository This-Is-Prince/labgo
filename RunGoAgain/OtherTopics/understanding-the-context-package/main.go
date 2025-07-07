package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var signal chan bool
var c = make(chan int)

func init() {
	signal = make(chan bool)
}

func printNumbers() {
	counter := 1
	for {
		select {
		case v, ok := <-signal:
			if ok {
				fmt.Println("Signal received:", v)
				return
			}
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println(counter)
			counter++
		}
	}
}

func square(ctx context.Context) {
	i := 0

	for {
		select {
		case <-ctx.Done():
			return
		case c <- i * i:
			i++
		}
	}
}

func main() {
	// go printNumbers()

	// fmt.Println("Num of Goroutines before signal:", runtime.NumGoroutine())
	// time.Sleep(1 * time.Second)

	// signal <- true

	// fmt.Println("Num of Goroutines after signal:", runtime.NumGoroutine())
	ctx, cancel := context.WithCancel(context.Background())

	go square(ctx)

	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	cancel()

	time.Sleep(3 * time.Second)

	fmt.Println("Number of Goroutines after cancel:", runtime.NumGoroutine())
}
