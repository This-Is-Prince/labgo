package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

/* ========== http.FileServer() function and http.Dir() type alias ========== */
/*
func main() {
	// create file server handler
	fs := http.FileServer(http.Dir("./tmp"))

	// start HTTP Server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":9000", fs))
}
*/

/* ========== Serving files on a route ========== */
/*
func main() {
	// create file server handler
	fs := http.FileServer(http.Dir("./tmp"))

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Golang!</h1>")
	})

	// handle `/static` route
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))
}
*/

/* ========== ServeFile function ========== */
/*
// temporary directory location
const cwd = "D:/Languages/Go/learning-go/RunGo/14-Beginers_guid_to_serving_files_using_HTTP_servers_in_Go"

var tmpDir = filepath.FromSlash(cwd + "/tmp/")

func main() {
	// return a `.pdf` file for `/pdf` route
	http.HandleFunc("/pdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, filepath.Join(tmpDir, "/files/test.pdf"))
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))
}
*/

/* ========== Handling index.html ========== */

// temporary directory location
const cwd = "D:/Languages/Go/learning-go/RunGo/14-Beginers_guid_to_serving_files_using_HTTP_servers_in_Go"

var tmpDir = filepath.FromSlash(cwd + "/tmp/")

func main() {
	// default route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// return a `.html` file for `/index.html` route
	http.HandleFunc("/index.html", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, filepath.Join(cwd, "/index.html"))
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))
}
