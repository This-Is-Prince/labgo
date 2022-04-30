package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling urls")

	fmt.Println(myurl)

	// parsing
	result, err := url.Parse(myurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("Port:", result.Port())
	fmt.Println("Path:", result.Path)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("RawQuery:", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println("coursename:", qparams.Get("coursename"))
	fmt.Println("paymentid:", qparams.Get("paymentid"))
	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Printf("%v : %v\n", key, val)
	}
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "loc.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

	response, err := http.Get(anotherURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	dataByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataByte))
}
