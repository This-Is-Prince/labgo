package main

import "fmt"

func main() {
	fmt.Println("The anatomy of Slices in Go!")

	var s []int
	fmt.Println(s == nil)
	fmt.Println(s, len(s), cap(s))

	s = []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = a[2:4]
	fmt.Println(s == nil, s)
	a[3] = 100
	fmt.Println(s, len(s), cap(s))

	// type T int
	// type T1 = uint32

	// t := T(10)

	fmt.Println(&a[2], &s[0])

	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a1[2:4]
	fmt.Println(s1, len(s1), cap(s1))
	s2 := append(s1, 100, 200, 300)
	fmt.Println(s2, len(s2), cap(s2))

	fmt.Println(&a1[2], &s1[0], &s2[0])
	s2[0] = -1
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}
