package ch1

import (
	"fmt"
	"os"
)

func Echo2() {
	var s, sep string
	for _, value := range os.Args[1:] {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
}
