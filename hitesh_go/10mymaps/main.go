package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["RS"] = "Rust"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("Key:%v | Value: %v\n", key, value)
	}
}
