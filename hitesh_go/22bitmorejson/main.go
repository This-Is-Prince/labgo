package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")

	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit321", nil},
	}

	// package this data as JSON data
	finalJSON, err := json.Marshal(lcoCourses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(finalJSON))

	finalIndentJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(finalIndentJSON))
}
