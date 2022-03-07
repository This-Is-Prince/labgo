package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("--------Web Request Verbs--------")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 1.first
	// var responseString strings.Builder
	// noOfBytes, err := responseString.Write(content)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("No of bytes returned from post request:", noOfBytes)
	// fmt.Println("Response return from post request:", responseString.String())

	// 2.second
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "prince")
	data.Add("lastname", "kumar")
	data.Add("email", "prince@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 1.first
	var contentString strings.Builder
	contentLength, err := contentString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("ContentLength is:", contentLength)
	fmt.Println("content is:", contentString.String())

	// 2.second
	fmt.Println("content is:", string(content))
}
