package main

import (
	// Global Package
	"fmt"
	"os"
	"strings"

	// Custom Package
	"files/src/file_creation"
	"files/src/file_reading"
	"files/src/file_writing"
)

func main() {
	fmt.Println("---------Handling Files In Golang---------")

	// HandlingFileCreations();
	HandlingFileReading()
}

func HandlingFileReading() {
	for _, filename := range os.Args[1:] {
		content := file_reading.ReadingFile(filename)
		fmt.Println(*content)
	}
}

func HandlingFileCreations() {
	content := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>File Creation And Writing</title>
  </head>
  <body>
    <h1 id="title"></h1>
  </body>
</html>
`
	for _, filename := range os.Args[1:] {
		CreateAndWrite(&filename, &content)
	}
}
func CreateAndWrite(filename *string, content *string) {
	file := file_creation.CreatingFile(*filename)
	defer file.Close()
	newContent := strings.Replace(*content, `<h1 id="title"></h1>`, fmt.Sprintf(`<h1 id="title">This %s file is Created, and Written By golang</h1>`, *filename), -1)
	file_writing.WritingFile(file, newContent)
}
