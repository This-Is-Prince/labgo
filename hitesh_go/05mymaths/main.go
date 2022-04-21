package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("-----------My Maths-----------")
	fmt.Println("Welcome to maths in golang")

	/* 	var myNumberOne int = 2
	   	var myNumberTwo float64 = 4
	   	fmt.Println("The sum is:", myNumberOne+int(myNumberTwo)) */

	// random number

	// through math
	/* 	rand.Seed(time.Now().UnixNano())
	   	fmt.Println(rand.Intn(6) + 1) */

	// through crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
