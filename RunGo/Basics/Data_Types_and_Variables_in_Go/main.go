package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Data Types And Variables in GO")

	// declaringAVariable()
	// typeInference()
	// shortHandNotation()
	// multipleVariableDeclarations()
	// moreOnShortHandNotation()
	// typeConversion()
	// definingAType()
	// constant()
	iotaExpression()
}

func declaringAVariable() {
	// var variableName dataType = initialValue

	var integer1 int = 15
	var integer2 int8 = -25
	var integer3 int32 // default 0
	var float1 float32 = 63.2
	var string1 string = "Hello World!"
	var boolean1 bool // default false
	var boolean2 bool = true

	// re-assign a value to variable
	// variableName = newValue
	integer1 = -15
	integer2 = 25
	integer3 = 965
	float1 = -52.99
	string1 = "Hello There :)"
	boolean1 = true
	boolean2 = false
	fmt.Println(integer1)
	fmt.Println(integer2)
	fmt.Println(integer3)
	fmt.Println(float1)
	fmt.Println(string1)
	fmt.Println(boolean1)
	fmt.Println(boolean2)
}

func typeInference() {
	// var variableName = initialValue

	var integer1 = 52            // int
	var string1 = "Hello World!" // string
	var boolean1 = false         // bool
	fmt.Println(integer1)
	fmt.Println(string1)
	fmt.Println(boolean1)
}

func shortHandNotation() {
	// variableName := initialValue

	integer1 := 52           // int
	string1 := "Hello World" // string
	boolean1 := false        // bool
	fmt.Println(integer1)
	fmt.Println(string1)
	fmt.Println(boolean1)
}

func multipleVariableDeclarations() {
	// var var1, var2, var3 int
	// var var1, var2, var3 int = 1, 2, 3
	// var var1, var2, var3 = 1, 2.2, false
	// var1, var2, var3 := 1, 2.2, false
	// var (
	// 	var1        = 50
	// 	var2        = 25.22
	// 	var3 string = "california"
	// 	var4 bool
	// )
}

func moreOnShortHandNotation() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("The minimum value is ", c)
}

func typeConversion() {
	var1 := 10   // int
	var2 := 10.5 // float64

	// illegal
	// var3 := var1 + var2

	// legal
	var3 := var1 + int(var2) // var3 => 20
	fmt.Println(var3)

	str1 := "Hello!"
	fmt.Println(str1)
	integer32 := []int32(str1)
	fmt.Println(integer32)
}

func definingAType() {
	derivedType()
	typeAlias()
}

type float float64

func derivedType() {
	// type newType fromType
	var f float = 52.2
	fmt.Printf("f has value %v and type %T\n", f, f)

	var f1 float = 52.2
	var f2 float64 = 52.2
	// fmt.Println("f1 == f2", f1 == f2) // ERROR, cannot compare f1 == f2 (mismatched types float and float64)
	fmt.Println("f1 == 52.2", f1 == 52.2) // Valid
	fmt.Println("f2 == 52.2", f2 == 52.2) // Valid
}

type float1 = float64

func typeAlias() {
	// type aliasName = oldType
	var f float1 = 52.2
	var g float64 = 52.2
	fmt.Println("f == g", f == g) // f == g true
}

func constant() {
	// const const_name data_type = fixed_value

	// const const_1, const_2 = value_1, value_2

	// const (
	// 	const_1 = value_1
	// 	const_2 = value_2
	// )

	const (
		a = 1 // a == 1
		b = 1 // b == 2
		c     // c == 2
		d     // d == 2
	)
}

func iotaExpression() {
	// const (
	// 	a = iota // a == 0
	// 	b = iota // b == 1
	// 	c = iota // c == 2
	// 	d        // d == 3 (implicitly d = iota)
	// )

	// const (
	// 	a = iota + 1 // a == 1
	// 	_            // (implicitly _ = iota + 1)
	// 	b            // b == 3 (implicitly b = iota + 1)
	// 	c            // c == 4 (implicitly c = iota + 1)
	// )

	// const a = iota // 0
	// const b = iota // 0

	const (
		c = iota // 0
		d = iota // 1
	)

	const (
		e = iota // 0
		f = iota // 1
	)
}
