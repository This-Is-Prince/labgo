package main

// Model for course - file
type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  string  `json:"price"`
	Author *Author `json:"author"`
}

//
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.Id == "" && c.Name == ""
}

func main() {

}
