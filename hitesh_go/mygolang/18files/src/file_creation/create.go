package file_creation

import (
	"os"

	"files/src/check_errors"
)

func CreatingFile(filename string) *os.File {
	file, err := os.Create(filename)
	check_errors.CheckNilError(err)
	return file
}
