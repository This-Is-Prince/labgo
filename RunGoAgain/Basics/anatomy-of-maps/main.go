package main

import "fmt"

func main() {
	fmt.Println("The anatomy of maps in Go!")

	var m map[string]int
	fmt.Println(m)
	fmt.Println("m == nil", m == nil)

	m1 := make(map[string]int)
	fmt.Println(m1)
	fmt.Println("m1 == nil", m1 == nil)

	age := make(map[string]int)
	age["prince"] = 24
	age["jasmeet"] = 25
	fmt.Println(age, len(age))
	fmt.Println(age["jasmeet"])
	jes, ok := age["jes"]
	fmt.Println(jes, ok)

	fmt.Println("Age:-", age)
	delete(age, "prince")
	fmt.Println("Age:-", age)

	for key, value := range age {
		fmt.Println("Key:", key, "Value:", value)
	}
}
