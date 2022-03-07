package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("-------Bit more JSON-------")

	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc1234", []string{"web-dev", "js"},
		},
		{
			"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd1234", []string{"full-stack", "js"},
		},
		{
			"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit1234", nil,
		},
	}

	// package this data JSON data
	// 1.first
	// finalJSON, err := json.Marshal(lcoCourses)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", finalJSON)

	// 2.second
	// finalJSON, err := json.MarshalIndent(lcoCourses, "P", "---")
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}
