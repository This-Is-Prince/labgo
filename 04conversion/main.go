package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 10")
	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	fmt.Printf("Rating is: %s", rating)
	fmt.Printf("Rating Type is: %T", rating)

	// numRating = input + 1 //error
	// numRating, err := strconv.ParseFloat(rating, 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Num Rating is: %f", numRating)
	// 	fmt.Printf("Num Rating Type is: %T", numRating)
	// }

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Num Rating is: %f", numRating)
		fmt.Printf("Num Rating Type is: %T", numRating)
	}
}
