package ch1

import (
	"fmt"
	"os"
)

func Echo1() {
	var sep, s string
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
