package main

import "fmt"

// jwtToken :=300000 // error this operator only allowed in function
const LoginToken string = "Prince Kumar ." // L capital public variable

func main() {
	fmt.Println(LoginToken)
	// var username string = "Prince"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T \n", username)

	// var isLoggedIn bool = true
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var smallValue uint8 = 255
	// fmt.Println(smallValue)
	// fmt.Printf("Variable is of type: %T \n", smallValue)

	// var smallFloat float64 = 255.45522354325
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "this.is.prince"
	fmt.Println(website)

	// no var style

	numberOfUser := 3000000
	fmt.Println(numberOfUser)
}
