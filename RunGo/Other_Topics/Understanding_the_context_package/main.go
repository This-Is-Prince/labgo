package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// func printNumbers() {
// 	counter := 1
// 	for {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(counter)
// 		counter++
// 	}
// }

// func first() {
// 	go printNumbers()
// 	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
// 	time.Sleep(time.Second)

// 	fmt.Println("After: active goroutines", runtime.NumGoroutine())
// 	fmt.Println("Program exited")
// }

// var cancelChannel = make(chan bool)

// func printNumbers2() {
// 	counter := 1

// 	for {
// 		select {
// 		case <-cancelChannel:
// 			return
// 		default:
// 			time.Sleep(100 * time.Millisecond)
// 			fmt.Println(counter)
// 			counter++
// 		}
// 	}
// }

// func second() {
// 	go printNumbers2()

// 	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
// 	time.Sleep(time.Second)

// 	cancelChannel <- true
// 	time.Sleep(1 * time.Millisecond)

// 	fmt.Println("After: active goroutines", runtime.NumGoroutine())
// 	fmt.Println("Program exited")
// }

// var c = make(chan int)

// func square(ctx context.Context) {
// 	i := 0
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case c <- i * i:
// 			i++
// 		}
// 	}
// }

// func third() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	fmt.Println("Number of active goroutines:", runtime.NumGoroutine())
// 	go square(ctx)
// 	fmt.Println("Number of active goroutines:", runtime.NumGoroutine())

// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Next square is:", <-c)
// 	}

// 	cancel()

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Number of active goroutines:", runtime.NumGoroutine())
// }

var startTime = time.Now()

func worker(ctx context.Context, seconds int) {
	select {
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), seconds)
		return
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Printf("%0.2fs - worker(%ds) completed the job.\n", time.Since(startTime).Seconds(), seconds)
	}
}

func fourth() {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	go worker(ctx, 2)
	go worker(ctx, 6)
	go worker(ctx, 8)

	time.Sleep(5 * time.Second)

	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}

func main() {
	fmt.Println("Understanding The Context Package")

	// first
	// first()

	// second
	// second()

	// third
	// third()

	// fourth
	fourth()
}
