package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Server")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port: 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<ul><li>Name: %s</li><li>Address: %s</li></ul>", name, address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "<h1>Hello Bro!</h1>")
}
