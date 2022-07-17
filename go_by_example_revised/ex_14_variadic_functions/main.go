package main

import "fmt"

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	nums := []int{1, 2, 3}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, value := range nums {
		total += value
	}
	fmt.Println(total)
}
