package main

import "fmt"

func main() {
	fmt.Println("Methods")

	prince := User{"Prince", "prince@go.dev", true, 21}
	fmt.Println(prince)

	prince.GetStatus()
	fmt.Println(prince)
	prince.NewMail()
	fmt.Println(prince)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() {
	fmt.Println("Is user active:", user.Status)
}

// func (u User) NewMail() {
// 	u.Email = "test@go.dev"
// 	fmt.Println("Email of this user is:", u.Email)
// }

func (u *User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
