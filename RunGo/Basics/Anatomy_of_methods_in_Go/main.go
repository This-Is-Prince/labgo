package main

import "fmt"

/* ========== What is Method? ========== */
/* type Employee struct {
	FirstName, LastName string
	FullName            func(string, string) string
}

func fullName(firstName, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func main() {
	e := Employee{
		FirstName: "Go",
		LastName:  "lang", FullName: func(firstName, lastName string) string {
			return firstName + " " + lastName
		},
	}
	fmt.Println(fullName(e.FirstName, e.LastName))
	fmt.Println(e.FullName(e.FirstName, e.LastName))
	fmt.Println(e.fullName())
}
*/

/* ========== Methods with the same name ========== */
/*
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
	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect= %0.2f\n", rect.Area())
	fmt.Printf("Area of Circle cir= %0.2f\n", cir.Area())
}
*/

/* ========== Pointer receivers ========== */
/*
type Employee struct {
	name   string
	salary int
}

func (e Employee) changename(newName string) {
	e.name = newName
}

func main() {
	e := Employee{
		name:   "Employee Di silva",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)

	// change name
	e.changename("NoBody")

	// e after name change
	fmt.Println("e after name change =", e)
}
*/

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changename(newName string) {
	// both are same
	(*e).name = newName
	// e.name = newName
}

func main() {
	e := Employee{
		name:   "Employee Di silva",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)

	// create pointer to `e`
	ep := &e
	// change name
	ep.changename("NoBody")

	// e after name change
	fmt.Println("e after name change =", e)
}
