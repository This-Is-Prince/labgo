package file_writing

import (
	"files/src/check_errors"
	"io"
	"os"
)

func WritingFile(file *os.File, content string) int {
	noOfBytesWrite, err := io.WriteString(file, content)
	check_errors.CheckNilError(err)
	return noOfBytesWrite
}
