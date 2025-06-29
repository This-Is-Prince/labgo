package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1!"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2!"
}

func main() {
	fmt.Println("Anatomy of Channels in GO - Concurrency in Go!")

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

// func greet(c chan string) {
// 	fmt.Println("Hello " + <-c + "!")
// }

// func greet1(c chan string) {
// 	<-c
// 	<-c
// }

// func squares(c chan int) {
// 	for i := 0; i <= 9; i++ {
// 		c <- i * i
// 	}

// 	close(c)
// }

// func squares1(c chan int) {
// 	for i := 0; i <= 3; i++ {
// 		num := <-c
// 		fmt.Println(num * num)
// 	}
// }

// func main() {
// 	fmt.Println("Anatomy of Channels in Go - Concurrency in Go!")

// 	// var c chan int
// 	// fmt.Println(c)

// 	c := make(chan int)

// 	fmt.Println(len(c), cap(c))

// 	fmt.Printf("type of c is %T\n", c)
// 	fmt.Printf("value of c is %v\n", c)

// 	c1 := make(chan string)
// 	go greet(c1)

// 	c1 <- "John"

// 	// c2 := make(chan string)

// 	// go greet1(c2)
// 	// c2 <- "John"

// 	// close(c2)

// 	// c2 <- "Mike"

// 	c3 := make(chan int)

// 	go squares(c3)

// 	for {
// 		val, ok := <-c3
// 		if !ok {
// 			fmt.Println(val, ok, "<-- loop broke!")
// 			break
// 		} else {
// 			fmt.Println(val, ok)
// 		}
// 	}

// 	c4 := make(chan int, 3)

// 	go squares1(c4)

// 	fmt.Println("first", len(c4), cap(c4))
// 	c4 <- 1
// 	fmt.Println("second", len(c4), cap(c4))
// 	c4 <- 2
// 	fmt.Println("third", len(c4), cap(c4))
// 	c4 <- 3
// 	fmt.Println("fourth", len(c4), cap(c4))
// 	c4 <- 4
// 	fmt.Println("last", len(c4), cap(c4))

// 	c5 := make(chan int, 3)

// 	go sender(c5)

// 	fmt.Printf("length of channel c is %v and capacity of channel c is %v\n", len(c5), cap(c5))

// 	for val := range c5 {
// 		fmt.Printf("length of channel c after value %v read is %v\n", val, len(c5))
// 	}

// 	roc := make(<-chan int)
// 	soc := make(chan<- int)

// 	fmt.Printf("Data type of roc is %T\n", roc)
// 	fmt.Printf("Data type of soc is %T\n", soc)

// 	c6 := make(chan string)

// 	go greet2(c6)

// 	c6 <- "Alice"

// 	go func(c chan string) {
// 		fmt.Println("Hello " + <-c + " from anonymous function!")
// 	}(c6)

// 	c6 <- "Bob"

// 	cc := make(chan chan string)

// 	go greeter(cc)

// 	c7 := <-cc

// 	go greet2(c7)

// 	c7 <- "Charlie"

// 	fmt.Println("main() stopped")
// }

// func sender(c chan int) {
// 	c <- 1
// 	c <- 2
// 	c <- 3
// 	c <- 4
// 	close(c)
// }

// func greet2(roc <-chan string) {
// 	fmt.Println("Hello " + <-roc + "!")
// }

// func greeter(cc chan chan string) {
// 	c := make(chan string)
// 	cc <- c
// }
