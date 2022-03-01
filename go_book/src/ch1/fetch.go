package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	const Prefix = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, Prefix) {
			url = Prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		// b, err := ioutil.ReadAll(resp.Body)
		written, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s\n", b)
		fmt.Printf("%d\n", written)
	}
}
