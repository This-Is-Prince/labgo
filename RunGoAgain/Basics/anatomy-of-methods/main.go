package main

import (
	"fmt"
	"math"
)

type Employee struct {
	FirstName, LastName string
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) ChangeFirstName(newName string) {
	e.FirstName = newName
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	fmt.Println("Anatomy of methods in Go!")

	e := Employee{
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Println(e.FullName())

	rect := Rect{5.0, 4.0}
	cir := Circle{3.0}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Area())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.Area())

	fmt.Println("Before:-", e)
	ep := &e
	ep.ChangeFirstName("Jane")
	fmt.Println("After:-", e)
	e.ChangeFirstName("Alice")
	fmt.Println("After changing again:-", e)
}
