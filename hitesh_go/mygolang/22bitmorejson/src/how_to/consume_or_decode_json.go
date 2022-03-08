package how_to

import (
	"bitmorejson/src/errors_util"
	"encoding/json"
	"fmt"
)

func Consume_Or_Decode_JSON() {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "ReactJS Bootcamp",
        "price": 299,
        "website": "LearnCodeOnline.in",
        "tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	isJsonValid := json.Valid(jsonDataFromWeb)
	if isJsonValid {
		fmt.Println("JSON was valid")
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		errors_util.CheckNilError(err)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	fmt.Println("-----------")

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	errors_util.CheckNilError(err)

	fmt.Println("-----------")
	for k, v := range myOnlineData {
		fmt.Printf("Key is `%v` and value is `%v` and type of value is `%T`\n", k, v, v)
	}
}
