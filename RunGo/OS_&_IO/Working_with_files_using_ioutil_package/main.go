package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
====================
ALL FUNCTIONS

1). ioutil.ReadDir(dirname string)
2). os.Stat(filepath string)
3). ioutil.ReadFile(filepath string)
4). ioutil.WriteFile(filepath string, data []byte, perm os.FileMode)

====================
*/

func main() {
	// ReadDirFunction()
	// ReadDirUsingGlobFunction()
	// ReadFileFunction()
	// WriteFileFunction()
	MkDirAllFunction()
}

func ReadDirFunction() {
	files, _ := ioutil.ReadDir("./tmp")
	for _, file := range files {
		fmt.Printf(
			"Name: %v, Size: %vkb\n",
			file.Name(),
			file.Size()/1024,
		)
		fmt.Printf("Mode: %v, IsDir: %v, IsRegular: %v, Perm: %v, String: %v, Type: %v\n",
			file.Mode(),
			file.Mode().IsDir(),
			file.Mode().IsRegular(),
			file.Mode().Perm(),
			file.Mode().String(),
			file.Mode().Type(),
		)
		fmt.Printf("%#v\n", file.Sys())
	}
}

func ReadDirUsingGlobFunction() {
	htmlFilesPath, _ := filepath.Glob("./tmp/**/*.html")
	htmlFilesPath1, _ := filepath.Glob("./tmp/*.html")
	htmlFilesPath = append(htmlFilesPath, htmlFilesPath1...)

	for _, htmlFilePath := range htmlFilesPath {
		file, _ := os.Stat(htmlFilePath)
		fmt.Printf(
			"Name: %v, Size: %vkb\n",
			file.Name(),
			file.Size()/1024,
		)
		fmt.Printf("Mode: %v, IsDir: %v, IsRegular: %v, Perm: %v, String: %v, Type: %v\n",
			file.Mode(),
			file.Mode().IsDir(),
			file.Mode().IsRegular(),
			file.Mode().Perm(),
			file.Mode().String(),
			file.Mode().Type(),
		)
	}
}

func ReadFileFunction() {
	bytes, err := ioutil.ReadFile("./tmp/style.css")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

func WriteFileFunction() {
	err := ioutil.WriteFile("./tmp/main.js", []byte(
		`console.log("hello")
console.log("hello")
`), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func MkDirAllFunction() {
	err := os.MkdirAll("./tmp/htmlFiles/home", 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./tmp/htmlFiles/home/index.html", []byte(
		`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <h1>
      This file is created with the help of os.MkdirAll() function and ioutil.WriteFile() function
    </h1>
  </body>
</html>
`), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
