package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello")

	fmt.Fprint(w, " World! ")

	data := []byte(`❤️`)

	w.Write(data)
}

func main() {
	// handler := HttpHandler{}

	// http.ListenAndServe(":9000", handler)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	http.ListenAndServe(":9000", mux)
}
