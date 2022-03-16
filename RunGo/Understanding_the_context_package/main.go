package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/* ====== Problem ====== */
/*
func printNumbers(){
	counter:=1

	for{
		time.Sleep(100*time.Millisecond)
		counter++
	}
}
func main() {
	go printNumbers()

	fmt.Println("Before: active goroutines",runtime.NumGoroutine())
	time.Sleep(time.Second)

	fmt.Println("After: active goroutines",runtime.NumGoroutine())
	fmt.Println("Program exited")
}
*/

/* ====== Stop Signal ====== */
/*
var signal = make(chan bool)

func printNumbers() {
	counter := 1
	for {
		select {
		case <-signal:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			counter++
		}
	}
}
func main() {
	go printNumbers()

	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
	time.Sleep(time.Second)

	signal <- true

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")
}
*/

/* ====== context.WithCancel ====== */
/*
// channel to send square of integers
var c = make(chan int)

// send square of numbers
func square() {
	i := 0
	for {
		i++
		c <- i * i
	}
}

// main goroutine
func main() {
	go square() // start square goroutine

	// get 5 square
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}
	// do other job
	time.Sleep(3 * time.Second)

	// print active goroutine
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}
*/
/*
// channel to send square of integers
var c = make(chan int)

// send square of numbers
func square(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return // kill goroutine
		case c <- i * i:
			i++
		}
	}
}

// main goroutine
func main() {
	// create cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	go square(ctx) // start square goroutine

	// get 5 square
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	fmt.Println("Before:Number of active goroutines", runtime.NumGoroutine())
	// cancel context
	cancel() // instead of `defer cancel()`
	// do other job
	time.Sleep(3 * time.Second)

	// print active goroutine
	fmt.Println("After:Number of active goroutines", runtime.NumGoroutine())
}
*/

/* ====== context.WithDeadline ====== */
/*
// start time
var startTime = time.Now()

// perform a job in the future
func worker(ctx context.Context, seconds int) {
	select {
	// if context closes, end goroutine
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), seconds)
		return // kills goroutine
		// do the job after `seconds` seconds
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Printf("%0.2fs - worker(%ds) completed the job.\n", time.Since(startTime).Seconds(), seconds)
	}
}

// main goroutine
func main() {
	// deadline => at time, 3 seconds in the future
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	// cancel if `main` returns before the deadline
	defer cancel()

	// start worker goroutines
	go worker(ctx, 2) // execute a job after 2 seconds
	go worker(ctx, 6) // execute a job after 6 seconds
	go worker(ctx, 8) // execute a job after 8 seconds

	// sleep for 5 seconds
	time.Sleep(5 * time.Second)

	// number of active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}
*/

/* ====== context.WithTimeout ====== */

// start time
var startTime = time.Now()

// perform a job in the future
func worker(ctx context.Context, seconds int) {
	select {
	// if context closes, end goroutine
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), seconds)
		return // kills goroutine
		// do the job after `seconds` seconds
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Printf("%0.2fs - worker(%ds) completed the job.\n", time.Since(startTime).Seconds(), seconds)
	}
}

// main goroutine
func main() {
	// deadline => at time, 3 seconds in the future
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))

	// cancel if `main` returns before the deadline
	defer cancel()

	// start worker goroutines
	go worker(ctx, 2) // execute a job after 2 seconds
	go worker(ctx, 6) // execute a job after 6 seconds
	go worker(ctx, 8) // execute a job after 8 seconds

	// sleep for 5 seconds
	time.Sleep(5 * time.Second)

	// number of active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}
