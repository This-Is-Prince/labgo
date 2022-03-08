package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"webreqverbs/src/errors_utils"
)

func PostForm() {
	const myURL = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "tyrion")
	data.Add("lastname", "lannister")
	data.Add("age", "47")
	data.Add("email", "tyrion@casterly.rock")

	response, err := http.PostForm(myURL, data)
	errors_utils.CheckNilError(err)
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	errors_utils.CheckNilError(err)

	var responseString strings.Builder
	countByte, err := responseString.Write(responseBytes)
	fmt.Println("No of bytes written in responseString:-", countByte)
	fmt.Println("Response String:-", responseString.String())
}
