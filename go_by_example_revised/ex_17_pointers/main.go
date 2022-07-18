package main

import "fmt"

func zeroVal(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}

func main() {
	i, j := 10, 20
	fmt.Println(i, j)
	zeroVal(i)
	zeroPtr(&j)
	fmt.Println(i, j)
	fmt.Println(&j)
}
