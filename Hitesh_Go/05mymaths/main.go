package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang...")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5
	// fmt.Println("The Sum is: ", myNumberOne+int(myNumberTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println( rand.Intn(5))

	// crypto random number

	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(randomNumber)
}
