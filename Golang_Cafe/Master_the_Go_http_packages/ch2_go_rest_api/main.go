package main

import (
	"fmt"
	"log"
	"net/http"
)

type userHandler struct{}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	fmt.Println("GO REST API")

	mux := http.NewServeMux()
	mux.Handle("/users", &userHandler{})
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
