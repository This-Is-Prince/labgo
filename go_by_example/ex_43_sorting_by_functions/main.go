package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Age  int
	Name string
}
type byAge []Person

func (p byAge) Len() int {
	return len(p)
}
func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fmt.Println("Sorting By Functions")

	fruits := []string{"peach", "banana", "kiwi"}
	fmt.Printf("%v isSorted? %v\n", fruits, sort.IsSorted(byLength(fruits)))
	sort.Sort(byLength(fruits))
	fmt.Printf("%v isSorted? %v\n", fruits, sort.IsSorted(byLength(fruits)))

	persons := []Person{{Age: 22, Name: "Prince"}, {Age: 21, Name: "Prince Kumar"}, {Age: 15, Name: "Kumar"}}

	fmt.Printf("%v isSorted? %v\n", persons, sort.IsSorted(byAge(persons)))
	sort.Sort(byAge(persons))
	fmt.Printf("%v isSorted? %v\n", persons, sort.IsSorted(byAge(persons)))
}
