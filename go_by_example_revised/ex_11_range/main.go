package main

import "fmt"

func main() {
	fmt.Println("Range")

	nums := []int{2, 3, 5}
	fmt.Println(nums)
	for _, v := range nums {
		fmt.Println(v)
	}

	maps := make(map[string]int)
	maps["first"] = 1
	maps["second"] = 2
	maps["third"] = 3
	for key, value := range maps {
		fmt.Println(key, ":", value)
	}

	str := "Hello World"
	for _, character := range str {
		fmt.Println(string(character))
	}
}
