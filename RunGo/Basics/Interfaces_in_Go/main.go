package main

import "fmt"

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
	var s Shape = Cube{3}
	c, found := s.(Cube)
	if !found {
		fmt.Println("Cube type not implement Shape interface")
		return
	}
	fmt.Println("area of c of type Cube is", c.Area())
	fmt.Println("volume of c of type Cube is", c.Volume())
}
*/

/* ========= Is Underlying value of an interface implements any other interfaces? ========= */
/*
type Shape interface {
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

func main() {
	var s Shape = Cube{3}
	value1, ok := s.(Object)
	fmt.Printf("dynamic value of Shape `s` with value %v implements interface Object? %v\n", value1, ok)
	value2, ok := s.(Skin)
	fmt.Printf("dynamic value of Shape `s` with value %v implements interface Sin? %v\n", value2, ok)
}
*/

/* ========= Type Switch ========= */
/*
func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}
func main() {
	explain("Hello World")
	explain(52)
	explain(true)
}
*/

/* ========= Embedding Interfaces ========= */
/*
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Material interface {
	Shape
	Object
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
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("dynamic type and value of interface `m` of static type Material is `%T` and `%v`\n", m, m)
	fmt.Printf("dynamic type and value of interface `s` of static type Shape is `%T` and `%v`\n", s, s)
	fmt.Printf("dynamic type and value of interface `o` of static type Object is `%T` and `%v`\n", o, o)
}
*/

/* ========= Pointer VS Value receiver ========= */
/*
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{5.0, 4.0}
	var s Shape = r // ERROR
	area := s.Area()
	fmt.Println("area of rectangle is", area)
}
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{5.0, 4.0}
	var s Shape = &r // assigned pointer
	area := s.Area()
	perimeter := s.Perimeter()
	fmt.Println("area of rectangle is", area)
	fmt.Println("perimater of rectangle is", perimeter)
}
