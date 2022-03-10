package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg *sync.WaitGroup

func main() {
	wg = &sync.WaitGroup{}
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		signals = append(signals, endpoint)
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

/* -----------------2.Second----------------- */
/*
import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg, rootWG *sync.WaitGroup

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	rootWG = &sync.WaitGroup{}
	rootWG.Add(1)
	go withGoRoutines(websitelist)
	rootWG.Wait()
	fmt.Println("------------------------")
	rootWG.Add(1)
	go withoutGoRoutines(websitelist)
	rootWG.Wait()
}

func withGoRoutines(websitelist []string) {
	defer rootWG.Done()
	start := time.Now()
	wg = &sync.WaitGroup{}
	for _, web := range websitelist {
		go getStatusCodeForWithGoRoutines(web)
		wg.Add(1)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time taken by these websites with go routines:-", elapsed)
}
func withoutGoRoutines(websitelist []string) {
	defer rootWG.Done()
	start := time.Now()
	for _, web := range websitelist {
		getStatusCodeForWithoutGoRoutines(web)
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken by these websites without go routines:-", elapsed)
}

func getStatusCodeForWithGoRoutines(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

func getStatusCodeForWithoutGoRoutines(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
} */

/* -----------------1.First----------------- */
/* import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------Go Routines--------")

	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		// time.Sleep(3 * time.Second)
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
*/
