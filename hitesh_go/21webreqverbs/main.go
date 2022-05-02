package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")

	// PerformGetRequest()
	PerformJsonRequest()
}

func PerformJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learnCodeOnline.in"	
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Request.GetBody()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"

	response, err := http.Get(myURL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	var builder strings.Builder

	lengthOfContent, err := builder.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Length of Content written in builder is:", lengthOfContent)
	fmt.Println(builder.String())
}
