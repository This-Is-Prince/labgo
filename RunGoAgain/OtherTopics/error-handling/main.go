package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "My custom error occurred"
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can not divide by zero")
	} else {
		return (a / b), nil
	}
}

type HttpError struct {
	status int
	method string
}

func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.", httpError.method, httpError.status)
}

func GetUserEmail(userId int) (string, error) {

	return "", &HttpError{403, "GET"}
}

func main() {
	myErr := &MyError{}

	// fmt.Println(myErr)
	fmt.Printf("Type of myErr is %T\n", myErr)
	fmt.Printf("Value of myErr is %#v\n", myErr)
	fmt.Println()

	newErr := errors.New("this is a new error")

	// fmt.Println(newErr)
	fmt.Printf("Type of newErr is %T\n", newErr)
	fmt.Printf("Value of newErr is %#v\n", newErr)
	fmt.Println()

	if result, err := Divide(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	if result, err := Divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	if email, err := GetUserEmail(1); err != nil {
		fmt.Println(err)

		if errVal, ok := err.(*HttpError); ok {
			fmt.Printf("Error is of type HttpError with status %d and method %s\n", errVal.status, errVal.method)
		} else {
			fmt.Println("Error is not of type HttpError")
		}
	} else {
		fmt.Println("User email:", email)
	}
	fmt.Println()
}
