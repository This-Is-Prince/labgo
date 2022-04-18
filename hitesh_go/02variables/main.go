package main

import "fmt"

// This is not allowed, outside any method
// jwtToken:=300000 // ERROR, expected declaration, found jwtToken
var jwtToken = 3000000

const LoginToken string = "abcdefgh"

func main() {
	fmt.Println("Variables")

	fmt.Println("\n---------- Global Variables ----------")
	fmt.Println(jwtToken)
	fmt.Printf("Variable `jwtToken` is of type: %T \n", jwtToken)

	fmt.Println("\n---------- Const Variables ----------")
	fmt.Println(LoginToken)
	fmt.Printf("Variable `LoginToken` is of type: %T \n", LoginToken)

	fmt.Println("\n---------- String Type ----------")
	var username string = "Prince"
	fmt.Println(username)
	fmt.Printf("Variable `username` is of type: %T \n", username)

	fmt.Println("\n---------- Boolean Type ----------")
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable `isLoggedIn` is of type: %T \n", isLoggedIn)

	fmt.Println("\n---------- Unsigned Integer Type ----------")
	// var smallVal uint8 = 256 // ERROR, cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable `smallVal` is of type: %T \n", smallVal)

	fmt.Println("\n---------- Float32 Type ----------")
	var float32Type float32 = 255.4554451125445188502384172842034827394871302142748273840273489702837428347023842703
	var float64Type float64 = 255.4554451125445188502384172842034827394871302142748273840273489702837428347023842703
	fmt.Println(float32Type)
	fmt.Println(float64Type)
	fmt.Printf("Variable `float32Type` is of type: %T \n", float32Type)
	fmt.Printf("Variable `float64Type` is of type: %T \n", float64Type)

	// default values and some aliases
	fmt.Println("\n---------- Default Values and Some Aliases ----------")
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable `anotherVariable` is of type: %T \n", anotherVariable)

	// implicit type
	fmt.Println("\n---------- Implicit Type (type inference) ----------")
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable `website` is of type: %T \n", website)

	// no var style
	fmt.Println("\n---------- no var style ----------")
	numberOfUser := 300000.
	fmt.Println(numberOfUser)
	fmt.Printf("Variable `numberOfUser` is of type: %T \n", numberOfUser)
}
