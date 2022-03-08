package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"webreqverbs/src/errors_utils"
)

func PostJson() {
	const myURL = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"Tyrion",
			"age":47,
			"married":true,
			"other-name":["Dwarf Of Casterly Rock","Beast","The Halfman","The Imp","The Little Monster"]
		}
	`)
	response, err := http.Post(myURL, "application/json", requestBody)
	errors_utils.CheckNilError(err)
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	errors_utils.CheckNilError(err)

	var responseString strings.Builder
	countByte, err := responseString.Write(responseBytes)
	errors_utils.CheckNilError(err)

	fmt.Println("No of bytes written in strings.Builder:-", countByte)
	fmt.Println("Response String:-", responseString.String())
}
