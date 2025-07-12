package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FirstName, lastName string
	Age                 int
	Email               string
	HeightInMeters      float64
	IsMale              bool
}

type Student1 map[string]interface{}

type Profile struct {
	Username  string
	Followers int
}

func (p Profile) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`{"f_count": %d}`, p.Followers)
	return []byte(str), nil
}

type Age int

func (a Age) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`{"age": %d}`, int(a))
	return []byte(str), nil
}

func (a Age) MarshalText() ([]byte, error) {
	str := fmt.Sprintf(`{"age": %d}`, int(a))
	return []byte(str), nil
}

type Student2 struct {
	FirstName, lastName string
	Age                 Age
	Profile             Profile
}

type Profile1 struct {
	Username  string `json:"uname"`
	Followers int    `json:"followers,omitempty,string"`
}

type Student3 struct {
	FirstName string   `json:"fname"`
	LastName  string   `json:"lname,omitempty"`
	Email     string   `json:"-"`
	Age       int      `json:"-,"`
	IsMale    bool     `json:",string"`
	Profile   Profile1 `json:""`
}

func main() {

	john := Student{
		FirstName:      "John",
		lastName:       "Doe",
		Age:            20,
		HeightInMeters: 1.75,
		IsMale:         true,
	}

	john1 := Student1{
		"FirstName":      "John",
		"lastName":       "Doe",
		"Age":            20,
		"HeightInMeters": 1.75,
		"IsMale":         true,
	}

	johnJSON, _ := json.Marshal(john)

	john1JSON, _ := json.Marshal(john1)

	fmt.Println(string(johnJSON))
	fmt.Println()

	fmt.Println(string(john1JSON))
	fmt.Println()

	john2 := &Student2{
		FirstName: "John",
		lastName:  "Doe",
		Age:       20,
		Profile: Profile{
			Username:  "johndoe",
			Followers: 100,
		},
	}

	john2JSON, _ := json.MarshalIndent(john2, "", "  ")
	fmt.Println(string(john2JSON))
	fmt.Println()

	john3 := &Student3{
		FirstName: "John",
		LastName:  "",
		Age:       21,
		Email:     "john@doe.com",
		Profile: Profile1{
			Username:  "johndoe91",
			Followers: 1975,
		},
	}

	john3JSON, _ := json.MarshalIndent(john3, "", "  ")
	fmt.Println(string(john3JSON))
	fmt.Println()
}
