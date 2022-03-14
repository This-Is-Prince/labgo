package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/* ========= Encoding JSON ========= */
/*
// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Email               string
	Age                 int
	HeightInMeters      float64
	IsMale              bool
}

func main() {
	// define `john` struct
	john := Student{
		FirstName:      "John",
		lastName:       "Doe",
		Age:            21,
		HeightInMeters: 1.75,
		IsMale:         true,
	}

	// encode `john` as JSON
	// johnJSON, err := json.Marshal(john)
	johnJSON, err := json.MarshalIndent(john, "", "\t")

	if err != nil {
		log.Fatal(err)
		return
	}

	// print JSON string
	fmt.Println(string(johnJSON))
}
*/

// Student declares `Student` map
/* type Student map[string]interface{}

func main() {
	// define `john` struct
	john := Student{
		"FirstName":      "John",
		"lastName":       "Doe",
		"Age":            21,
		"HeightInMeters": 1.75,
		"IsMale":         true,
	}

	// encode `john` as JSON
	// johnJSON, err := json.Marshal(john)
	johnJSON, err := json.MarshalIndent(john, "", "\t")

	if err != nil {
		log.Fatal(err)
		return
	}

	// print JSON string
	fmt.Println(string(johnJSON))
}
*/

/* ========= Data Types Handling ========= */
/* ========= Abstract Data Types ========= */
/*
// Profile declares `Profile` structure
type Profile struct {
	Username  string
	followers int
	Grades    map[string]string
}

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Age                 int
	Profile             Profile
	Languages           []string
}

func main() {
	var john Student

	// define `john` struct
	john = Student{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile{
			Username:  "johndoe007",
			followers: 1975,
			Grades:    map[string]string{"Math": "A", "Science": "A+"},
		},
		Languages: []string{"English", "French"},
	}

	// encode `john` as JSON
	johnJSON, err := json.MarshalIndent(john, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}

	// print JSON string
	fmt.Println(string(johnJSON))
} */

// anonymously nested structure
/*
// Profile declares `Profile` structure
type Profile struct {
	Username  string
	followers int
	Grades    map[string]string
}

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Age                 int
	Languages           []string
	Profile
}

func main() {
	var john Student

	// define `john` struct
	john = Student{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile{
			Username:  "johndoe007",
			followers: 1975,
		},
		Languages: []string{"English", "French"},
	}

	// encode `john` as JSON
	johnJSON, err := json.MarshalIndent(john, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}

	// print JSON string
	fmt.Println(string(johnJSON))
}
*/

//  ProfileI interface defines `Follow` method
/*
type ProfileI interface {
	Follow()
}

// Profile declares `Profile` structure
type Profile struct {
	Username  string
	followers int
}

func (p *Profile) Follow() {
	p.followers++
}

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Age                 int
	Primary             ProfileI
	Secondary           ProfileI
}

func main() {

	// define `john` struct (pointer)
	john := &Student{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Primary: &Profile{
			Username:  "johndoe007",
			followers: 1975,
		},
	}

	// follow `john`
	john.Primary.Follow()

	// encode `john` as JSON
	johnJSON, err := json.MarshalIndent(john, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}

	// print JSON string
	fmt.Println(string(johnJSON))
}
*/

/* ========== Data Type Conversion ========== */

//  Profile declares `Profile` structure
type Profile struct {
	Username  string
	Followers int
}

// MarshalJSON - implement `Marshaler` interface
func (p Profile) MarshalJSON() ([]byte, error) {
	// return JSON value (TODO: handle error gracefully)
	return []byte(fmt.Sprintf(`{"f_count": "%d"}`, p.Followers)), nil
}

// Age declares `Age` type
type Age int

// MarshalText - implement `TextMarshaler` interface
func (a Age) MarshalText() ([]byte, error) {
	// return string value (TODO: handle error gracefully)
	return []byte(fmt.Sprintf(`{"age": %d}`, int(a))), nil
}

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Age                 Age
	Profile             Profile
}

func main() {
	// define `john` struct (pointer)
	john := &Student{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile{
			Username:  "johndoe007",
			Followers: 1975,
		},
	}

	// encode `john` as JSON
	johnJSON, err := json.MarshalIndent(john, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}
	// print JSON string
	fmt.Println(string(johnJSON))
}
