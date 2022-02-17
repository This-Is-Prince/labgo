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

	// how to remove a value from a slices based on index
	fmt.Println()
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Value of courses is ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Value of courses is ", courses)
}
