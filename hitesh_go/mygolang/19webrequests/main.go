package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("-----Handling Web Requests in golang-----")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of type: %T", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
