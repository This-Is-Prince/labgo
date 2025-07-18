package main

import "fmt"

const (
	a4 = 1
	b4 = 2
	c4
	d4
)

const (
	a5 = iota + 1
	_
	b5
	c5
	d5
)

const (
	a6 = iota << 1
	b6
	c6
	d6
)

type MyStruct struct{}

// This will give an error because complete struct is a new type so if field change struct type change.
// type Hi MyStruct {
// 	A string
// }

type NewEquivalentFloat32Type float32
type AliasFloat32Type = float32

func Another() {
	fmt.Println("......................Another Starting......................")

	fmt.Println(a4, b4, c4, d4)
	fmt.Println(a5, b5, c5, d5)
	fmt.Println(a6, b6, c6, d6)

	_fp := 23.333
	fp := NewEquivalentFloat32Type(23.333)
	fmt.Println("Is NewEquivalentFloat32Type equal ", NewEquivalentFloat32Type(_fp) == fp)

	_fp1 := float32(23.333)
	fp1 := AliasFloat32Type(23.333)
	fmt.Println("Is AliasFloat32Type equal ", _fp1 == fp1)

	var a string = "Hey"
	var b rune = 3
	var c byte = 2
	var d int = 22
	var e float32 = 3.2
	var f float64 = 22.222323
	var g bool = true

	fmt.Println(a, b, c, d, e, f, g)

	var a1, b1, c1, d1, e1, f1, g1 = "Hey", rune('3'), byte('2'), 22, float32(3.2), 22.22223, true

	fmt.Println(a1, b1, c1, d1, e1, f1, g1)

	a2, b2, c2, d2, e2, f2, g2 := "Hey", rune('3'), byte('2'), 22, float32(3.2), 22.22223, true

	fmt.Println(a2, b2, c2, d2, e2, f2, g2)

	var (
		a3 = []byte("Hey")
		b3 = []rune("Hey")
		c3 = rune('8')
		d3 = byte('2')
		e3 = float32(22.2)
		f3 = float64(22323.23423434)
		g3 = false
	)

	fmt.Println(a3, b3, c3, d3, e3, f3, g3)
}
