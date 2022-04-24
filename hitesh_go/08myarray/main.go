package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Fruit List is:", fruitList)
	fmt.Println("Fruit List length is:", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veg List is:", vegList)
	fmt.Println("Veg List length is:", len(vegList))
}
