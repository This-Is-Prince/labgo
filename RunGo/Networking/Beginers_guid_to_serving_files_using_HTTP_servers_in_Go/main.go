package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Beginners Guide to serving files using http servers")

	// serving_A_Directory()
	serving_Files_On_A_Route()
}

func serving_A_Directory() {
	// create file server handler
	currDir, err := os.Getwd()
	fmt.Println(currDir, err)
	dirPath := filepath.Join(currDir, "tmp")
	fmt.Println(dirPath)
	fs := http.FileServer(http.Dir(dirPath))

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":9000", fs))
}

func serving_Files_On_A_Route() {
	// create file server handler
	currDir, _ := os.Getwd()
	fs := http.FileServer(http.Dir(filepath.Join(currDir, "tmp")))

	// handler `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Golang!</h1>")
	})

	// handle `/static` route
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))
}
