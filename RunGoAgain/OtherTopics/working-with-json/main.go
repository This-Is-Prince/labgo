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

type Student4 struct {
	FirstName, lastName string
	Email               string
	Age                 int
	HeightInMeters      float64
}

type Profile2 struct {
	Username  string
	Followers int
}

type Student5 struct {
	FirstName, lastName string
	HeightInMeters      float64
	IsMale              bool
	Languages           [2]string
	Subjects            []string
	Grades              map[string]string
	// Profile2            Profile2
	Profile2
}

type Student6 struct {
	FirstName, lastName string
	HeightInMeters      float64
	IsMale              bool
	Languages           [2]string
	Subjects            []string `json:",omitempty"`
	Grades              map[string]string
	Profile             *Profile2
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

	data := []byte(`
	{
		"FirstName": "John",
		"lastName": "Doe",
		"Age": 20,
		"HeightInMeters": 1.75,
		"Username": "johndoe91"
	}
	`)

	var john4 Student4
	err := json.Unmarshal(data, &john4)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("%#v\n", john4)
	fmt.Println()

	data1 := []byte(`
	{
		"FirstName": "John",
		"HeightInMeters": 1.75,
		"IsMale": null,
		"Languages": ["English", "Spanish", "German"],
		"Subjects": ["Math", "Science"],
		"Grades": { "Math": "A", "Art": "B" },
		"Profile2": {
			"Username": "johndoe91",
			"Followers": 1975
		}
	}
	`)
	// data1 := []byte(`
	// {
	// 	"FirstName": "John",
	// 	"HeightInMeters": 1.75,
	// 	"IsMale": null,
	// 	"Languages": ["English", "Spanish", "German"],
	// 	"Subjects": ["Math", "Science"],
	// 	"Grades": { "Math": "A", "Art": "B" },
	// 	"Username": "johndoe91",
	// 	"Followers": 1975
	// }
	// `)

	var john5 Student5 = Student5{
		IsMale:    true,
		Languages: [2]string{"Russian"},
		Subjects:  []string{"Art"},
		Grades:    map[string]string{"Science": "A+"},
	}

	fmt.Printf("Error: %v\n", json.Unmarshal(data1, &john5))
	fmt.Printf("%#v\n", john5)
	fmt.Println()

	data2 := []byte(`
	{
		"FirstName": "John",
		"HeightInMeters": 1.75,
		"IsMale": null,
		"Languages": ["English", "Spanish", "German"],
		"Subjects": null,
		"Grades": null,
		"Profile": { "Followers": 1975 }
	}
	`)

	var john6 Student6 = Student6{
		IsMale:    true,
		Languages: [2]string{"Russian"},
		Subjects:  []string{"Art"},
		Grades:    map[string]string{"Math": "A"},
		Profile:   &Profile2{Username: "johndoe91"},
	}

	fmt.Printf("Error: %v\n", json.Unmarshal(data2, &john6))
	fmt.Printf("%#v\n", john6)
	fmt.Printf("%#v\n", john6.Profile)
	fmt.Println()
}
