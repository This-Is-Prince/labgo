package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sorting")

	strings := []string{"c", "a", "b"}
	fmt.Printf("%v isSorted? %v\n", strings, sort.StringsAreSorted(strings))
	sort.Strings(strings)
	fmt.Printf("%v isSorted? %v\n", strings, sort.StringsAreSorted(strings))

	ints := []int{7, 2, 4}
	fmt.Printf("%v isSorted? %v\n", ints, sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Printf("%v isSorted? %v\n", ints, sort.IntsAreSorted(ints))
}
