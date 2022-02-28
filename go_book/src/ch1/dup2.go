package ch1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Counts struct {
	filename string
	data     map[string]int
}

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
func Dup2() {
	files := os.Args[1:]
	n := len(files)
	countsArr := make([]Counts, n)
	for i := 0; i < n; i++ {
		countsArr[i] = Counts{data: make(map[string]int)}
	}
	if n == 0 {
		countsArr[0] = Counts{filename: "stdin", data: make(map[string]int)}
		countLines(os.Stdin, countsArr[0])
	} else {
		for index, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsArr[index].filename = filepath.Base(f.Name())
			countLines(f, countsArr[index])
			f.Close()
		}
	}
	for _, counts := range countsArr {
		fmt.Println("---------FileName:", counts.filename, "---------")
		for line, n := range counts.data {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts Counts) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts.data[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
}
