package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Length of Veg list is: ", len(vegList))
}
