package unsafepackage

import (
	"fmt"
	"unsafe"
)

func Learn(run bool) {
	if !run {
		return
	}

	s := "Hello, World!"
	s1 := 'H'
	fmt.Printf("String: %s, Type: %T\n", s, s)
	fmt.Printf("Rune: %c, Type: %T\n", s1, s1)

	for i, v := range s {
		fmt.Printf("%c %v %T, %c %v %T\n", s[i], s[i], s[i], v, v, v)
	}

	fmt.Println("Learning unsafe package")

	number := int32(100)
	pointer := unsafe.Pointer(&number)
	fmt.Println("Pointer to number:", pointer, &number)

	// textPointer := (*string)(pointer)
	// fmt.Println("Pointer to text:", textPointer, *textPointer)
	// fmt.Println(number)

	pi := float64(3.14159)

	// bits := unsafe.Sizeof(pi)
	// fmt.Println("Size of pi in bytes:", bits)
	bits := *(*uint64)(unsafe.Pointer(&pi))
	fmt.Println("Bits representation of pi:", bits)

	bits1 := *(*float32)(unsafe.Pointer(&pi))
	fmt.Println("Bits representation of pi as float32:", bits1, pi)

	a := int32(10)
	fmt.Printf("Value of a: %d, Type: %T\n", a, a)
	b := *(*int16)(unsafe.Pointer(&a))
	fmt.Printf("Value of b: %d, Type: %T\n", b, b)

	c := 'A'
	fmt.Printf("Value of c: %c, Type: %T\n", c, c)
	d := *(*int32)(unsafe.Pointer(&c))
	fmt.Printf("Value of d: %d, Type: %T\n", d, d)

	e := "Hello"
	fmt.Printf("Value of e: %s, Type: %T\n", e, e)
	f := *(*[]int16)(unsafe.Pointer(&e))
	fmt.Printf("Value of f: %v, Type: %T\n", f, f)

	num := float32(12.34)
	floatKey := &num

	masterKey := unsafe.Pointer(floatKey)

	intKey := (*uint32)(masterKey)
	fmt.Printf("Original float: %f\n", *floatKey)
	fmt.Printf("Raw bits as integer: %d\n", *intKey)

	// The unsafe package is used for low-level programming in Go.
	// It allows you to manipulate memory directly, which can be dangerous.
	// Use it with caution and only when necessary.

	// Example usage of unsafe package would go here.
	// For instance, using unsafe.Pointer to convert between types,
	// or using unsafe.Sizeof to get the size of a type in bytes.
}
