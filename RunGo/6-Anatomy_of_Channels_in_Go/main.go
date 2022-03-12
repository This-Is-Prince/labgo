package main

import (
	"fmt"
	"sync"
)

/* =============== Declaring a channel =============== */
/*
func main() {
	// var c chan int
	// fmt.Println(c)

	c := make(chan int)
	fmt.Printf("type of 'c' is %T\n", c)
	fmt.Printf("value of 'c' is %v\n", c)
}
*/

/* =============== Channels in practice =============== */
/*
func greet(c chan string) {
	fmt.Println("Hello "+<-c, "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "Prince"
	fmt.Println("main() stopped")
}
*/

/* =============== Deadlock =============== */
/*
func main() {
	fmt.Println("main() started")

	c := make(chan string)
	c <- "Prince" // fatal error: all goroutines are asleep - deadlock!

	fmt.Println("main() stopped")
}
*/

/* =============== Closing a channel =============== */
/*
func greet(c chan string) {
	<-c // for Prince
	<-c // for Kumar
}

func main() {
	fmt.Println("main() started")

	c := make(chan string)

	go greet(c)

	c <- "Prince"

	close(c) // closing channel

	c <- "Kumar" // panic: send on closed channel
	fmt.Println("main() stopped")
}
*/

/* =============== For Loop =============== */
/*
func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until channel closes
	// for {
	// 	val, ok := <-c
	// 	if ok {
	// 		fmt.Println(val, ok)
	// 	} else {
	// 		fmt.Println(val, ok, "<-- loop broke!")
	// 		break
	// 	}
	// }

	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main() stopped")
}
*/

/* =============== Buffer size or channel capacity =============== */
/*
func squares(c chan int) {
	for i := 0; i <= 6; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here
	fmt.Println("Five")
	c <- 5
	time.Sleep(time.Millisecond)

	fmt.Println("main() stopped")
}
*/

/* =============== length and capacity of a channel =============== */
/*
func sender(c chan int) {
	fmt.Println("one")
	c <- 1 // len 1, cap 3
	fmt.Println("two")
	c <- 2 // len 2, cap 3
	fmt.Println("three")
	c <- 3 // len 3, cap 3
	fmt.Println("four")
	c <- 4 // goroutine blocks here
	close(c)
}

func main() {
	c := make(chan int, 3)
	// c <- 1
	// c <- 2

	// fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
	// fmt.Println()

	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	// read values from c (blocked here)
	for val := range c {
		fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
	}
}
*/

/* =============== brain teaser =============== */
/*
func squares(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)
	go squares(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())

	go squares(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())

	c <- 5
	c <- 6
	c <- 7
	c <- 8 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")
}
*/

/* =============== reading data from closed channels, data lives in the buffer =============== */
/*
func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// iteration terminates after receiving 3 values
	for elem := range c {
		fmt.Println(elem)
	}
}
*/

/* =============== Working with multiple goroutines =============== */
/*
func square(c chan int) {
	fmt.Println("[square] reading")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("[cube] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	fmt.Println("[main] sent testNum to squareChan")

	squareChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")

	cubeChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube of", testNum, "is", sum)
	fmt.Println("[main] main() stopped")
}
*/

/* =============== Unidirectional Channels =============== */
/*
func main() {
	receive_only_channel := make(<-chan int)
	send_only_channel := make(chan<- int)

	fmt.Printf("Data type of receive_only_channel is %T\n", receive_only_channel)
	fmt.Printf("Data type of send_only_channel is %T\n", send_only_channel)
}
*/

/* =============== convert bi-directional channel to unidirectional channel. =============== */
/*
func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
	// roc <- "Prince" // invalid operation: cannot send to receive-only channel roc (variable of type <-chan string)
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "Prince"
	fmt.Println("main() stopped")
}
*/

/* =============== Anonymous goroutine. =============== */
/*
func main() {
	fmt.Println("main() started")
	c := make(chan string)

	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "Prince"
	fmt.Println("main() stopped")
} */

/* =============== Channel as the data type of channel =============== */
/*
// gets a channel and prints the greeting by reading from channel
func greet(c <-chan string) {
	fmt.Println("Hello " + <-c + "!")
}

// gets a channels and write a channel to it
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func main() {
	fmt.Println("main() started")

	// make a channel `cc` of data type channel of string dta type
	cc := make(chan chan string)

	go greeter(cc) // start `greeter` goroutine using `cc` channel

	// receive a channel `c` from `greeter` goroutine
	c := <-cc

	go greet(c) // start  `greet goroutine` using `c` channel

	// send data to `c` channel
	c <- "Prince"

	fmt.Println("main() stopped")
}
*/

/* =============== Select =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	// time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	// time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}
*/

/* =============== To simulate when all the cases are non-blocking (buffered channel) =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}
*/

/* =============== `default` case =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	default:
		fmt.Println("No response received", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}
*/

/* =============== `default` case =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	time.Sleep(3 * time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	default:
		fmt.Println("No response received", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
} */

/* =============== `default` case =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	default:
		fmt.Println("No goroutine available to send data", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}
*/

/* =============== nil channel =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}
func service(c chan string) {
	c <- "response"
}

func main() {
	fmt.Println("main() started")
	var chan1 chan string
	go service(chan1)

	select {
	case res := <-chan1: // ERROR, select (no cases) means that select statement is virtually empty because cases with nil channel are ignored. but as empty select{} statemet blocks the main goroutine and service goroutine is scheduled in its place, channel operation on nil channels throws chan send (nil chan) error.
		fmt.Println("Response from chan1", res)
	}
	fmt.Println("main() stopped")
} */

/* =============== nil channel =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}
func service(c chan string) {
	c <- "response"
}

func main() {
	fmt.Println("main() started")
	var chan1 chan string
	go service(chan1)

	select {
	case res := <-chan1: // no error, default block executed
		fmt.Println("Response from chan1", res)
	default:
		fmt.Println("No response")
	}
	fmt.Println("main() stopped")
}
*/

/* =============== Adding timeout =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}
func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}
func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service1(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(3 * time.Second):
		fmt.Println("No response received", time.Since(start))

	}
	fmt.Println("main() stopped", time.Since(start))
}
*/

/* =============== Empty select =============== */
/*
func service() {
	fmt.Println("Hello from service")
}
func main() {
	fmt.Println("main() started")

	go service()
	select {}

	fmt.Println("main() stopped") // Unreachable code
}
*/

/* =============== WaitGroup =============== */
/*
func service(wg *sync.WaitGroup, instance int) {
	defer wg.Done() // decrement counter
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
}

func main() {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter
		go service(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}
*/

/* =============== Worker Pool =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

// worker than make squares
func sqlWorker(tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}

func main() {
	fmt.Println("[main] main() started", time.Since(start))

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqlWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receiving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped", time.Since(start))
}
*/

/* =============== Worker Pool With WaitGroup =============== */
/*
var start time.Time

func init() {
	start = time.Now()
}

// worker than make squares
func sqlWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
	wg.Done()
}

func main() {
	fmt.Println("[main] main() started", time.Since(start))

	wg := &sync.WaitGroup{}
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqlWorker(wg, tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	wg.Wait()

	// receiving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped", time.Since(start))
}
*/

/* =============== Mutex =============== */
/*
var i int // i==0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup) {
	i = i + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	// wait until all 1000 goroutines are done
	wg.Wait()

	// value of i should be 1000
	fmt.Println("value of i after 1000 operations is", i) // i will be less than 1000 or equal to 1000 everything happens
}
*/

var i int // i==0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(wg, mut)
	}

	// wait until all 1000 goroutines are done
	wg.Wait()

	// value of i should be 1000
	fmt.Println("value of i after 1000 operations is", i) // i will be less than 1000 or equal to 1000 everything happens
}
