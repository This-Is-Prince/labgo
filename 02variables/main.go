package main

import "fmt"

// jwtToken := 300000.0 // This is not allowed
// var jwtToken = 300000.0 //This is Allowed

const LoginToken string = "Prince Kumar"; //First Character is Capital 'L' 

func main()  {
	var username string="Prince";
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool=true;
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var hi int =2000;
	// fmt.Println(hi)
	// fmt.Printf("Variable is of type: %T \n", hi)
	
	// var smallValue uint8=256; // Error
	var smallValue uint8=255;
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)
	
		// var smallFloat float32=255.532534523452345234;
		// fmt.Println(smallFloat)
		// fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64=255.532534523452345234;
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// Default Values and some aliases

	var anotherVariable int;
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
	
	// Implicit Type or implicit way

	var website="prince.com";
	// website=3; // Error
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)
	
	fmt.Println(LoginToken);
	fmt.Printf("Variable is of type: %T \n", LoginToken)
	
}