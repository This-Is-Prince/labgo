package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	prince := User{"Prince", "prince@go.dev", true, 21}
	fmt.Println(prince)
	fmt.Printf("prince details are: %+v\n", prince)
	fmt.Printf("Name is %v and email is %v.\n", prince.Name, prince.Email)
	prince.GetStatus()
	prince.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", prince.Name, prince.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
