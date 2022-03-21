package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("Running Multiple HTTP Server In Go")

	// launch_Another_Goroutine()
	// launch_Two_Goroutine()
	// custom_Server()
	// close_Method()
	shutdown_Method()
}

func launch_Another_Goroutine() {
	// create a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: "+req.Host)
	})

	// create a goroutine
	go func() {
		// spawn a HTTP server in `other` goroutine
		log.Fatal(http.ListenAndServe(":9000", nil))
	}()

	// spawn an HTTP server in `main` goroutine
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func launch_Two_Goroutine() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two groutines to `wg` WaitGroup
	wg.Add(2)

	// create a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: "+req.Host)
	})

	// goroutine to launch a server on port 9000
	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
		wg.Done()
	}()

	// goroutine to launch a server on port 9001
	go func() {
		log.Fatal(http.ListenAndServe(":9001", nil))
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
}

func createServer(name string, port int) *http.Server {
	// create `ServeMux`
	mux := http.NewServeMux()

	// create a default route handler
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: "+name)
	})

	// create new server
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port), // :{port}
		Handler: mux,
	}

	// return new server (pointer)
	return &server
}

func custom_Server() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// goroutine to launch a server on port 9000
	go func() {
		server := createServer("ONE", 9000)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	// goroutine to launch a server on port 9001
	go func() {
		server := createServer("TWO", 9001)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
}

func close_Method() {
	// create channels for safe exit
	exitSignal := make(chan interface{})

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `DefaultServeMux`
	}

	// close server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Close(): completed!", server.Close()) // close `server`
		close(exitSignal)                                  // close `exitSignal` channel
	})

	// launch server
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)

	// listen to `exitSignal` channel
	<-exitSignal
	fmt.Println("Main(): exit complete!")
}

func shutdown_Method() {
	// create a `WaitGroup` for safe exit
	wg := new(sync.WaitGroup)
	wg.Add(2) // add `2` goroutines to finish

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `DefaultServeMux`
	}

	// register a cleanup function
	server.RegisterOnShutdown(func() {
		fmt.Println("RegisterOnShutdown(): completed!") // perfect a cleanup
		wg.Done()                                       // WaitGroup done
	})

	// shutdown server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		err := server.Shutdown(context.Background()) // shutdown `server`
		fmt.Println("Shutdown(): completed!", err)
		wg.Done() // WaitGroup done
	})

	// launch server
	fmt.Println("ListenAndServe():", server.ListenAndServe())

	// listen for `wg` to complete
	wg.Wait()
	fmt.Println("Main(): exit complete")
}
