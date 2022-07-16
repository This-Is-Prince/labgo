package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Functions")

	first, second, n := split("Hello World Bro")
	fmt.Println(first, second, n)
}

func split(str string) (string, string, int) {
	arr := strings.Split(str, " ")
	return arr[0], strings.Join(arr[1:], " "), len(arr[1:])
}
