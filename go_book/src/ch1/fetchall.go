package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const HTTP_PREFIX = "http://"
const HTTPS_PREFIX = "https://"

func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(i, url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(i int, url string, ch chan<- string) {
	if !strings.HasPrefix(url, HTTP_PREFIX) && !strings.HasPrefix(url, HTTPS_PREFIX) {
		url = HTTPS_PREFIX + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	_, fileName, found := strings.Cut(url, "//")
	pathname := "./test/fetchall/"
	if found {
		fileName = pathname + strings.ReplaceAll(fileName, "/", "_") + ".html"
	} else {
		fileName = pathname + fmt.Sprint("temp%d.html", i)
	}
	file, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprintf("while Creating file %s: %v", fileName, err)
	}
	defer file.Close()
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
