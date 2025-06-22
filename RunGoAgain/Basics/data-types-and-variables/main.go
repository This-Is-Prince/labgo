package main

import "fmt"

const CONST_NAME = "Hello, World!"
const CONST_NAME1 int = 100
const v1, v2, v3, v4 = 1, "Hey", false, 3.14

const (
	a = 1
	b = 2
	c
	d
)

const (
	e = iota
	f = iota
	g = iota
	h
	i
)

const (
	ee = iota + 1
	ff = iota
	gg = iota
	hh
	ii
)

const (
	eee = iota + 1
	_
	fff = iota
	ggg = iota
	hhh
	iii
)

func main() {
	fmt.Println("Constant:", CONST_NAME)
	fmt.Println("Constant with type:", CONST_NAME1)
	fmt.Println("Multiple constants:", v1, v2, v3, v4)
	fmt.Println("Constants with iota:", e, f, g, h, i)
	fmt.Println("Constants with iota and offset:", ee, ff, gg, hh, ii)
	fmt.Println("Constants with iota and skip:", eee, fff, ggg, hhh, iii)
	fmt.Println(a, b, c, d)
	fmt.Println("Hello, World!")

	// Declare a variable
	// var variableName dataType = initialValue
	var age int = 30
	fmt.Println("Age:", age)
	var age1 int
	fmt.Println("Age1:", age1)

	var integer1 int = 15
	var integer2 int8 = -25
	var integer3 int32 // default 0
	var float1 float32 = 63.2
	var string1 string = "Hello, World!"
	var boolean1 bool // default false
	var boolean2 bool = true

	fmt.Println(integer1, integer2, integer3, float1, string1, boolean1, boolean2)

	integer1 = -15
	integer2 = 25
	integer3 = 965
	float1 = 2.36
	string1 = "Hello, Go!"
	boolean1 = true
	boolean2 = false
	fmt.Println(integer1, integer2, integer3, float1, string1, boolean1, boolean2)

	// Declare variable without type
	// var variableName = initialValue
	var integer4 = 52
	var string2 = "Hello, Go Again!"
	var boolean3 = false

	fmt.Println(integer4, string2, boolean3)

	// Declare variable with shorthand syntax
	// variableName := initialValue
	integer5 := 100
	string3 := "Shorthand Syntax"
	boolean4 := true
	fmt.Println(integer5, string3, boolean4)

	// Multiple variable declaration
	// var var1, var2, var3 dataType
	// var var1, var2, var3 dataType = initialValue1, initialValue2, initialValue3
	var var1, var2, var3 int
	var var4, var5, var6 = "go", false, 3.14
	var7, var8, var9 := float64(22), -32, uint8(2)

	fmt.Println(var1, var2, var3, var4, var5, var6, var7, var8, var9)

	var (
		a = 1
		b = -21
		c = "Multiple declaration"
		d = false
	)

	fmt.Println(a, b, c, d)

	var11 := 10
	var22 := 10.5

	var33 := var11 + int(var22)
	fmt.Println("Sum of var11 and var22:", var33)

	h := "Hello, Go!"
	hInt32 := []int32(h) // int32 is rune type in Go
	fmt.Println("String to int32 slice:", hInt32)

	hUint8 := []uint8(h) // uint8 is byte type in Go
	fmt.Println("String to uint8 slice:", hUint8)

	// Derived Type
	// type newType fromType

	type Float float64
	float := Float(3.14)
	var float2 Float = 2.14
	var float3 Float
	fmt.Println("Float type:", float, float2, float3)

	fmt.Printf("float has value %v and type %T\n", float, float)

	var f Float = 52.2
	var g float64 = 52.2
	// fmt.Println("f:", f, "g:", g, "f == g:", f == g) // invalid operation f == g (mismatched types Float and float64
	fmt.Println("f:", f, "g:", g, "f == g:", f == Float(g))

	// Type Alias
	// type aliasName = oldType
	type MyInt = int
	myInt := MyInt(42)
	fmt.Println("MyInt value:", myInt)
	onlyInt := 100
	fmt.Println("Only int value:", onlyInt)
	fmt.Println("MyInt == onlyInt:", myInt == onlyInt) // true, because MyInt is an alias for int
}
