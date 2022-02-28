package ch1

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		if err := input.Err(); err != nil {
			fmt.Println(err)
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
