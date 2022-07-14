package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3}
	fmt.Println(arr)

	var strArr [5]string
	fmt.Println(strArr)
	strArr[0] = "Zero"
	strArr[2] = "Two"
	strArr[4] = "Four"
	fmt.Println(strArr)
	fmt.Println("Len:", len(strArr))

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
