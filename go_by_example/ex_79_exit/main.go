package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("EXIT")

	defer fmt.Println("!")

	os.Exit(3)
}
