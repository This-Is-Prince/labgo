package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"webreqverbs/src/errors_utils"
)

func Get() {
	const myURL = "http://localhost:8000/get"

	response, err := http.Get(myURL)
	errors_utils.CheckNilError(err)
	defer response.Body.Close()

	fmt.Println("Status:-", response.Status)
	fmt.Println("StatusCode:-", response.StatusCode)
	fmt.Println("ContentLength:-", response.ContentLength)

	contentInBytes, err := ioutil.ReadAll(response.Body)
	errors_utils.CheckNilError(err)

	var contentInString strings.Builder
	countByte, err := contentInString.Write(contentInBytes)
	errors_utils.CheckNilError(err)

	fmt.Println("Total Bytes Written:-", countByte)
	fmt.Println("Content String Get:-", contentInString.String())

}
