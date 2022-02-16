package main

import "fmt"

// jwtToken:=300000 // This(syntax) is not allowed outside any methods
var jwtToken int = 300000

const LoginToken string = "abcdefghijklmnop"

func main() {
	/*
		Variables
	*/
	var username string = "Prince"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "this_is_prince.in"
	fmt.Println(website)
	website = "prince.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 10000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)
}
