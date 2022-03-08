package file_reading

import (
	"files/src/check_errors"
	"io/ioutil"
	"strings"
)

func ReadingFile(filename string) *string {
	databytes, err := ioutil.ReadFile(filename)
	check_errors.CheckNilError(err)

	// Two Ways to convert databytes to string
	// 1. Using string()
	// content := string(databytes)
	// return &content

	// 2. Using strings.Builder

	var builder strings.Builder
	_, err = builder.Write(databytes)
	check_errors.CheckNilError(err)

	content := builder.String()
	return &content
}
