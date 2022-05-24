package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	fmt.Println("GO Routines")

	websiteList := []string{"https://lco.dev", "https://go.dev", "https://google.com", "https://fb.com", "https://github.com"}

	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}
	wg.Wait()

	mut.Lock()
	fmt.Println(signals)
	mut.Unlock()
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
	}
}

// First GO ROutines
/*
func main() {
	fmt.Println("Go Routines")

	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
*/
