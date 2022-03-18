package main

/*
import (
	"fmt"
	"log"
	"net/http"
) */

/* ============== The ListenAndServe function ============== */
/*
type HttpHandler struct {
}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
		// create response binary data
	   	// data := []byte(`Hello World!`) // slice of bytes

	   	// write `data` to response
	   	// res.Write(data)

	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello")

	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " World! ")

	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
}

func main() {
	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":9000", handler)
}
*/

/* ============== Understanding ServeMux ============== */
/*
func main() {
	// create a new `ServeMux`
	mux := http.NewServeMux()

	// handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve using `ServeMux`
	http.ListenAndServe(":9000", mux)
}
*/

/* ============== Using http.DefaultServeMux ============== */
/*
func main() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve using `http.DefaultServeMux`
	http.ListenAndServe(":9000", nil)
}
*/

/* ============== Returning a better response ============== */
/*
func main() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2020")

		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400

		// respond with a JSON string
		fmt.Fprint(res, `{"status": "FAILURE"}`)
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))
}
*/
