package main

import "fmt"

func main() {
	fmt.Println("---------For---------")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		// i = i + 1
		// i += 1
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
