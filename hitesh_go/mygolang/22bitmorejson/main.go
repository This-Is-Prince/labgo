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

	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
            "coursename": "ReactJS Bootcamp",
            "price": 299,
            "website": "LearnCodeOnline.in",
            "tags": ["web-dev","js"]
        }
	`)

	var lcoCourse course
	fmt.Printf("%#v\n", lcoCourse)

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println(("JSON WAS NOT VALID"))
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type of value is %T\n", k, v, v)
	}
}
