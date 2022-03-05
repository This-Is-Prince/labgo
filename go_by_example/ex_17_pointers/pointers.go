package main

import "fmt"

func zeroval(i_val int) {
	i_val = 0
}

func zeroptr(i_ptr *int) {
	*i_ptr = 0
}

func main() {
	fmt.Println("-----------Pointers-----------")

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
