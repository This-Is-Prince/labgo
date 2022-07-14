package main

import "fmt"

func main() {
	fmt.Println("Slices")

	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s = append(s, "d")
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(len(c), cap(c))

	withoutFirst := c[1:]
	fmt.Println(withoutFirst)
	fmt.Println(len(withoutFirst), cap(withoutFirst))
}
