package how_to

import (
	"bitmorejson/src/errors_util"
	"encoding/json"
	"fmt"
	"strings"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func Create_Or_Encode_JSON() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// Ways to make json
	// 1. Without Proper Indentation
	// finaljsonInBytes, err := json.Marshal(lcoCourses)

	// 2. With Proper Indentation
	finaljsonInBytes, err := json.MarshalIndent(lcoCourses, "", "\t")
	errors_util.CheckNilError(err)

	var finaljsonInString strings.Builder
	byteCount, err := finaljsonInString.Write(finaljsonInBytes)

	fmt.Println("No of bytes written in strings.Builder:-", byteCount)
	fmt.Println("Final JSON in string:-\n", finaljsonInString.String())
}
