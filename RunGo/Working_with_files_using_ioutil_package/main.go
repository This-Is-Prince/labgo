package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// temporary directory

func main() {
	fmt.Println("--------Working with files using `ioutil` package--------")

	// 1. ReadDir function
	/*
	   const tmpDir = "/tmp"
	   	cwd, err := os.Getwd()
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	}
	   	// get files from `/tmp` directory

	   	files, err := ioutil.ReadDir(path.Join(cwd, tmpDir))
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	}

	   	// list information of each file
	   	for _, file := range files {
	   		fmt.Printf("Name: %v, Size: %v kb, Mode: %v, IsDir: %v\n",
	   			file.Name(),
	   			file.Size()/1000,
	   			file.Mode(),
	   			file.IsDir(),
	   		)
	   	}
	*/

	//  2. Glob pattern
	// search for `.pdf` files
	// const tmpDir = "/tmp"
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// pdfFilesPathPattern := path.Join(cwd, tmpDir, "**/*.pdf")
	// fmt.Println("pdf files path pattern:", pdfFilesPathPattern)
	// pdfFiles, err := filepath.Glob(pdfFilesPathPattern)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// // list information of each pdf file
	// for _, pdfFile := range pdfFiles {
	// 	file, err := os.Stat(pdfFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	fmt.Printf("Name: %v, Size: %v kb, Mode: %v, IsDir: %v\n",
	// 		file.Name(),
	// 		file.Size()/1000,
	// 		file.Mode(),
	// 		file.IsDir(),
	// 	)
	// }
	// htmlFilesPathPattern := path.Join(cwd, tmpDir, "*.html")
	// fmt.Println("html files path pattern:", htmlFilesPathPattern)
	// htmlFiles, err := filepath.Glob(htmlFilesPathPattern)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// // list information of each html file
	// for _, htmlFile := range htmlFiles {
	// 	file, err := os.Lstat(htmlFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	fmt.Printf("Name: %v, Size: %v kb, Mode: %v, IsDir: %v\n",
	// 		file.Name(),
	// 		file.Size()/1000,
	// 		file.Mode(),
	// 		file.IsDir(),
	// 	)
	// }

	// 3. ReadFile function
	// location of index.html file
	/*
	   html, err := ioutil.ReadFile("./tmp/index.html")
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	}

	   	// print raw bytes
	   	fmt.Println("Raw: \n", html)

	   	// print string representation
	   	fmt.Println("String: \n", string(html))
	   	fmt.Printf("String: \n%s\n", html)

	   	// convert bytes to string
	   	var builder strings.Builder
	   	noOfBytesWrite, err := builder.Write(html)
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	}
	   	fmt.Println("No of bytes written:", noOfBytesWrite)
	   	fmt.Println("String: \n", builder.String())
	*/

	// 4. WriteFile function

	// welcome message content
	welcomeData := []byte("Welcome to my website.")

	// get `welcome.txt` file path
	welcomeFilePath := filepath.Join("./tmp", "/welcome.txt")

	// write content to `welcome.txt` file
	err := ioutil.WriteFile(welcomeFilePath, welcomeData, 0777)

	// log error
	if err != nil {
		log.Fatal(err)
	}
}
