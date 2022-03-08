package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs,angularjs&paymentid=ghbj45ghb,abcsghbj4"

func main() {
	fmt.Println("----------Handling URLs in Golang----------")

	// Parsing A URL
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:-", result.Scheme)
	fmt.Println("Hosname:-", result.Hostname())
	fmt.Println("Port:-", result.Port())
	fmt.Println("Host:-", result.Host)
	fmt.Println("Path:-", result.Path)
	fmt.Println("RawQuery:-", result.RawQuery)
	fmt.Println("User:-", result.User)
	fmt.Println("RawFragment:-", result.RawFragment)
	fmt.Println("Opaque:-", result.Opaque)
	fmt.Println("Fragment:-", result.Fragment)
	fmt.Println("ForceQuery:-", result.ForceQuery)

	fmt.Println("---------Query Parameters---------")
	qparams := result.Query()
	fmt.Println("Coursename:-", qparams["coursename"])
	fmt.Println("Password:-", qparams["password"])
	fmt.Println("PaymentId:-", qparams["paymentid"])

	fmt.Println("---------Query Parameters For Loop---------")
	for key, value := range qparams {
		fmt.Println(key, ":", value)
	}

	// Constructing A URL
	partsOfURL := &url.URL{
		Scheme:  "https",
		Path:    "/tutcss",
		Host:    "lco.dev",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
