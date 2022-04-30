package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("LCO web requests")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // caller's responsibility to close the connection
	fmt.Printf("Response is of type: %T\n", response)

	bytedata, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytedata)
	fmt.Println(content)
}
