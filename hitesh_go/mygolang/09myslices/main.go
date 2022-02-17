package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fmt.Println("Fruit list is: ", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit list is: ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("Fruit list is: ", fruitList)

	highScores := make([]int, 4)
	fmt.Println("Value of highScores is: ", highScores)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 //error, panic
	highScores = append(highScores, 555, 666, 321)
	fmt.Println("Value of highScores is: ", highScores)
	fmt.Println("Is highScores slices sorted? ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Value of highScores is: ", highScores)
	fmt.Println("Is highScores slices sorted? ", sort.IntsAreSorted(highScores))
}
