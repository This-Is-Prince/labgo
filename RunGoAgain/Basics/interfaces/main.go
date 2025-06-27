package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Interfaces in Go!")

	var s Shape

	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectangle s", s.Area())
	fmt.Println("s == r is", s == r)

	var s1 Shape = Rect{10, 3}
	fmt.Printf("type of s1 is %T\n", s1)
	fmt.Printf("value of s1 is %v\n", s1)
	fmt.Printf("value of s1 is %0.2f\n\n", s.Area())

	s1 = Circle{10}
	fmt.Printf("type of s1 is %T\n", s1)
	fmt.Printf("value of s1 is %v\n", s1)
	fmt.Printf("value of s1 is %0.2f\n\n", s.Area())

	ms := MyString("Hello World!")
	r1 := Rect{5.5, 4.5}
	explain(ms)
	explain(r1)

	ss, ok := s.(Circle)
	fmt.Println(ss, ok)

	var s2 Shape1 = Cube{3}
	value1, ok1 := s2.(Object)
	fmt.Printf("dynamic value of SHape 's2' with value %v implements interface Object? %v\n", value1, ok1)
	value2, ok2 := s2.(Skin)
	fmt.Printf("dynamic value of SHape 's2' with value %v implements interface Object? %v\n", value2, ok2)
}

type MyString string

func explain(i interface{}) {
	fmt.Printf("value given to explain functions is of type %T with value %v\n", i, i)
}

type Shape interface {
	Area() float64
	Perimeter() float64
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

func (r Rect) Perimeter() float64 {
	return 2 * r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape1 interface {
	Area() float64
}
type Object interface {
	Volume() float64
}
type Skin interface {
	Color() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * c.side * c.side
}
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
