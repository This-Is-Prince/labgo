package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP Clients")

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 500; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
