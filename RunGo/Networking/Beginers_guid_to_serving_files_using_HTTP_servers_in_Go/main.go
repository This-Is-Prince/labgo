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
	// serving_Files_On_A_Route()
	// serveFile_Function()
	handling_Index_HTML()
}

func serving_A_Directory() {
	// create file server handler
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := filepath.Join(currDir, "tmp")
	fmt.Println("DirPath:", dirPath)
	fs := http.FileServer(http.Dir(dirPath))

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":9000", fs))
}

func serving_Files_On_A_Route() {
	// create file server handler
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
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

func serveFile_Function() {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tmpDir := filepath.FromSlash(currDir)
	// return a `.pdf` file for `/pdf` route
	http.HandleFunc("/pdf", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.Path)
		http.ServeFile(res, req, filepath.Join(tmpDir+"/tmp/files/test.pdf"))
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handling_Index_HTML() {
	// working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// removing `slash` with platform dependant `\`, `/`
	currDir := filepath.FromSlash(wd)

	// default route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// return a `.html` file for `/index.html` route
	http.HandleFunc("/index.html", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, filepath.Join(currDir, "/index.html"))
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))
}
