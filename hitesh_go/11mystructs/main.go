package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	prince := User{"Prince", "prince@go.dev", true, 21}
	fmt.Println(prince)
	fmt.Printf("prince details are: %+v\n", prince)
	fmt.Printf("Name is %v and email is %v.", prince.Name, prince.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
