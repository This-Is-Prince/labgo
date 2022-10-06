package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("GO LoadBalancer")

	servers := NewSimpleServers(
		[]string{
			"https://www.facebook.com",
			"http://www.bing.com",
			"http://www.duckduckgo.com",
		})

	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(w http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(w, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}
