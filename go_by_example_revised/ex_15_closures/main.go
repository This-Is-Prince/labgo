package main

import (
	"fmt"
	"math"
	"strconv"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("Closures")

	anonymous := func() {
		fmt.Println("This is Anonymous Function")
	}
	anonymous()

	addTwoNumber := func(a, b int) (string, int) {
		return strconv.Itoa(int(math.Min(float64(a), float64(b)))), a + b
	}
	fmt.Println(addTwoNumber(1, 2))
	fmt.Println(addTwoNumber(5, 3))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	func(nums ...int) {
		total := 0
		for _, value := range nums {
			total += value
		}
		fmt.Println("Total:", total)
	}(1, 2, 3, 4, 5)
}
