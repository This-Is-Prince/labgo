package main

import "fmt"

func main() {
	fmt.Println("String Data Type In Go!")

	var s string

	s = "Hello World!"
	fmt.Println(s)
	fmt.Println("len is:- ", len(s))

	for i := 0; i < len(s); i++ {
		v := s[i]
		fmt.Printf("%v %c %T\n", v, v, v)
	}

	s = "Hellõ World!"
	fmt.Println(s)
	fmt.Println("len is:- ", len(s))

	for i := 0; i < len(s); i++ {
		v := s[i]
		fmt.Printf("%v %c %T\n", v, v, v)
	}

	s = "Hellõ World!"
	r := []rune(s)
	fmt.Println(r)
	fmt.Println("len is:- ", len(r))

	for i := 0; i < len(r); i++ {
		v := r[i]
		fmt.Printf("%v %c %T\n", v, v, v)
	}

	for idx, char := range s {
		fmt.Printf("Idx %v, Rune %v, Char %c\n", idx, char, char)
	}

	r1 := []rune("H")
	r2 := 'H'

	fmt.Printf("R1 %v %T, R2 %v %T", r1, r1, r2, r2)

}
