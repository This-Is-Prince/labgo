package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./tmp"))

	// log.Fatal(http.ListenAndServe(":9000", fs))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, `<h1>Golang!</h1> <a href="/static/hi.html">Hi</a>`)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./tmp/hi.html")
	})

	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":9000", nil))
}
