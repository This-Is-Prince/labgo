package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myURL = "https://lco.dev"

func main() {
	fmt.Println("-------Making Get Web Request-------")

	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// Conversion bytes to string
	var builder strings.Builder
	noOfBytesWritten, err := builder.Write(databytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("No of Bytes Written:", noOfBytesWritten)
	content := builder.String()
	fmt.Println("Length Of String:", len(content))
	fmt.Println(content)
}
