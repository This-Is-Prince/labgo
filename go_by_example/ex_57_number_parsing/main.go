package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Number Parsing")

	f, err := strconv.ParseFloat("1.234", 64)
	HandleError(err)
	fmt.Println(f)

	i, err := strconv.ParseInt("123", 0, 64)
	HandleError(err)
	fmt.Println(i)

	d, err := strconv.ParseInt("0x1c8", 0, 64)
	HandleError(err)
	fmt.Println(d)

	u, err := strconv.ParseUint("789", 0, 64)
	HandleError(err)
	fmt.Println(u)

	k, err := strconv.Atoi("135")
	HandleError(err)
	fmt.Println(k)

	s, err := strconv.Atoi("wat")
	HandleError(err)
	fmt.Println(s)
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
