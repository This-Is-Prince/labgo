package main

import (
	"fmt"
)

/* ========= Declaring Interface ========= */
/*
type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {
	var s Shape
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)
}
*/

/* ========= Implementing interface ========= */
/*
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
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// func main(){
// 	var s Shape
// 	s=Rect{5.0,4.0}
// 	r:=&Rect{5.0,4.0}
// 	fmt.Printf("type of s is %T\n",s)
// 	fmt.Printf("value of s is %v\n",s)
// 	fmt.Println("area of rectangle s",s.Area())
// 	fmt.Println("s == r is",s==r)
// }

func main() {
	var s Shape = Rect{10, 3}

	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("value of s is %0.2f\n\n", s.Area())

	s = Circle{10}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("value of s is %0.2f\n\n", s.Area())
}
*/

/* ========= Empty interface ========= */

/*
type MyString string

type Rect struct{
	width float64
	height float64
}

func explain(i interface{}){
	fmt.Printf("value given to explain function is of type `%T` with value %v\n",i,i)
}

func main(){
	ms:=MyString("Hello World!")
	r:=Rect{5.5,4.5}
	explain(ms)
	explain(r)
}
*/

/* ========= Multiple interface ========= */
/*
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
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

func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c
	fmt.Println("volume of s of interface type Shape is", s.Area())
	fmt.Println("area of o of interface type Object is", o.Volume())
}
*/

/* ========= Type assertion ========= */

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
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

func main() {
	var s Shape = Cube{3}
	c, found := s.(Cube)
	if !found {
		fmt.Println("Cube type not implement Shape interface")
		return
	}
	fmt.Println("area of c of type Cube is", c.Area())
	fmt.Println("volume of c of type Cube is", c.Volume())
}
