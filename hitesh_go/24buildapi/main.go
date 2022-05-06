package main

import "fmt"

// Model for course - file
type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.Id == "" && c.Name == ""
}

func main() {
	fmt.Println("Building API")
}
