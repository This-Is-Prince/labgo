package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go!")

	a := 0x00
	b := 0x0A
	c := 0xFF

	fmt.Println("&a =", &a)
	fmt.Println("&b =", &b)
	fmt.Println("&c =", &c)

	var pa *int
	fmt.Printf("Pointer pa of type %T with value %v\n", pa, pa)

	a1 := 1
	var pa1 *int
	pa1 = &a1
	fmt.Printf("Pointer pa1 of type %T with value %v\n", pa1, pa1)
	fmt.Printf("data at %v is %v\n", pa1, *pa1)

	*pa1 = 2
	fmt.Printf("Pointer pa1 of type %T with value %v\n", pa1, pa1)
	fmt.Printf("data at %v is %v\n", pa1, *pa1)

	pa2 := new(int)
	fmt.Printf("data at %v is %v\n", pa2, *pa2)

	p := 2
	fmt.Println(p)
	changeValue(&p)
	fmt.Println(p)

	arr := [3]int{4, 5, 6}
	fmt.Println(arr)
	changeValue1(&arr)
	fmt.Println(arr)
	changeValue2(&arr)
	fmt.Println(arr)
}

func changeValue(p *int) {
	*p = 20
}

func changeValue1(p *[3]int) {
	(*p)[0] = 1
	(*p)[1] = 2
	(*p)[2] = 3
}

func changeValue2(p *[3]int) {
	p[0] *= 8
	p[1] *= 9
	p[2] *= 10
}
