package main

import "fmt"

func main() {
	var a string
	fmt.Println(a, len(a))
	var b string = "aa"
	fmt.Println(b, len(b))

	var c = true
	fmt.Println(c)

	var d, e, f = 1, true, 4.9
	fmt.Println(d, e, f)
	var g, h, j uint8 = 255, 255, 255
	fmt.Println(g, h, j)

	hello := "Hello"
	fmt.Println(hello)

	i, k := true, 6.7
	fmt.Println(i, k)
}
