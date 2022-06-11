package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		fmt.Fprintf(w, "Next Headers\n")
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\nNext Header\n", name, h)
		}
	}
}

func main() {
	fmt.Println("HTTP Servers")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":4000", nil)
}
