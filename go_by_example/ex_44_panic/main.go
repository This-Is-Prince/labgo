package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Panic")

	// panic("a problem")

	_, err := os.Create("./tmp/index.html")
	if err != nil {
		panic(err)
	}
}
