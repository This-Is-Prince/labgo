package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8000", nil)
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	io.WriteString(writer, "Hello World!")
}
