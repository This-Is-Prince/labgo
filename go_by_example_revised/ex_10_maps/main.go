package main

import "fmt"

func main() {
	fmt.Println("Maps")

	m := make(map[string]int)
	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2

	fmt.Println(m)
	v, isAvailable := m["zer"]
	fmt.Println(v, isAvailable)
	v, isAvailable = m["zero"]
	fmt.Println(v, isAvailable)

	m1 := map[string]int{"zero": 0, "one": 1}
	fmt.Println(m1)

	fmt.Println(len(m))
	fmt.Println(len(m1))

	delete(m1, "zero")
	fmt.Println(len(m1))
}
