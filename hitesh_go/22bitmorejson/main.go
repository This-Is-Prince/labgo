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

	// EncodeJson()
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
                "coursename": "ReactJS Bootcamp",
                "price": 299,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`)

	var locCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		err := json.Unmarshal(jsonDataFromWeb, &locCourse)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", locCourse)
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
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
