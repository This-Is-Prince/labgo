package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main()  {
	fmt.Println("Welcome To Maths In GoLang.")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("The Sum is: ",myNumberTwo+float64(myNumberOne))
	// fmt.Println("The Sum is: ",myNumberOne+int(myNumberTwo))

	// random number
	// rand.Seed(30232)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))
	
	// random from crypto

	myRandomNumber , _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(myRandomNumber)
}