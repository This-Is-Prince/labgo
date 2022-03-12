package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/* =========== 1. Global constants and Variables =========== */

/* ===== 1. os.PathSeparator ===== */
/*
func main() {
	fmt.Printf("os.PathSeparator => %v\n", os.PathSeparator)
	fmt.Printf("os.PathSeparator => %c\n", os.PathSeparator)
}
*/

/* ===== 2. os.DevNull ===== */
/*
func main() {
	fmt.Printf("os.DevNull => %v\n", os.DevNull)
}
*/

/* ===== 3. os.Args ===== */
/*
func main() {
	fmt.Printf("os.Args => %v\n", os.Args)
}
*/

/* ===== 4. Exiting a process with os.Exit(code int) ===== */
// Status Code `0` indicates process was exited `Successfully`
// Status Code `1` indicates process was exited with a `general error`
/* func main() {
	fmt.Println("main() started!")

	// defer function execution
	defer func() {
		fmt.Println("main completed!")
	}()

	// run function after 2 seconds
	time.AfterFunc(2*time.Second, func() { os.Exit(1) }) // exit with status code : 1

	// sleep goroutine for 3 seconds
	time.Sleep(2 * time.Second)
	// hello will not printed
	fmt.Println("Hello")
} */

/* ===== 5. Exiting a process with process.Kill()  ===== */
/*
func main() {
	fmt.Println("main() started!")
	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		fmt.Println("defer function")
	}()

	process.Kill()
	fmt.Println("main() stopped!")
} */

/* =========== 2. Handling Environment Variables =========== */

/* ===== 1. os.Environ() ===== */
/*
func main() {
	// get all environment variables
	var env []string = os.Environ()

	// iterate over each variable
	for index, variable := range env {
		// get name and value of the variable
		nameVal := strings.Split(variable, "=")
		fmt.Printf("[%d]. %v => %v\n", index, nameVal[0], nameVal[1])
	}
}
*/

/* ===== 2. os.Getenv() and os.LookupEnv() ===== */
/*
func main() {
	// read `FRUIT` environment variable
	fruit := os.Getenv("FRUIT")
	fmt.Printf("os.Getenv(`FRUIT`) => %v\n", fruit)

	// read `CAR` environment variable (empty)
	car := os.Getenv("CAR")
	fmt.Printf("os.Getenv(`CAR`) => %v\n", car)

	// read `COUNTRY` environment variable (missing)
	country := os.Getenv("COUNTRY")
	fmt.Printf("os.Getenv(`COUNTRY`) => %v\n", country)

	//-------------

	// read `FRUIT` environment variable
	fruit, fruitExists := os.LookupEnv("FRUIT")
	fmt.Printf("os.LookupEnv(`FRUIT`) => %v (exists: %v)\n", fruit, fruitExists)

	// read `CAR` environment variable (empty)
	car, carExists := os.LookupEnv("CAR")
	fmt.Printf("os.LookupEnv(`CAR`) => %v (exists: %v)\n", car, carExists)

	// read `COUNTRY` environment variable (missing)
	country, countryExists := os.LookupEnv("COUNTRY")
	fmt.Printf("os.LookupEnv(`COUNTRY`) => %v (exists: %v)\n", country, countryExists)
}
*/

/* ===== 3. os.Setenv() and os.Unsetenv() ===== */
/*
func main() {
	// set environment variables
	fmt.Printf("Setenv(FRUIT) => error(%v)\n", os.Setenv("FRUIT", "banana")) // exists
	fmt.Printf("Setenv(CAR) => error(%v)\n\n", os.Setenv("CAR", "audi"))     // missing

	// print values
	fmt.Printf("Env: FRUIT/CAR => %v/%v\n\n", os.Getenv("FRUIT"), os.Getenv("CAR"))

	// unset environment variables
	fmt.Printf("Unsetenv(FRUIT) => error(%v)\n", os.Unsetenv("FRUIT"))       // exists
	fmt.Printf("Unsetenv(COUNTRY) => error(%v)\n\n", os.Unsetenv("COUNTRY")) // missing

	// print values
	fmt.Printf("Env: FRUIT/CAR => %v/%v\n\n", os.Getenv("FRUIT"), os.Getenv("CAR"))
}
*/

/* ===== 4. os.Expand() and os.ExpandEnv() ===== */
/*
func mapper(placeholder string) string {
	switch placeholder {
	case "FRUIT":
		return "mango"
	case "CAR":
		return "audi"
	default:
		return "<nil>"
	}
}
func main() {
	// raw string
	raw := "I am eating $FRUIT and driving ${CAR} and talking to $FRIEND"

	// formatted string
	formatted := os.Expand(raw, mapper)

	// print formatted string
	fmt.Println(formatted)
}
*/

/* ===== 5. os.Expand() and os.ExpandEnv() ===== */
/*
func main() {
	// set environment variable
	os.Setenv("COUNTRY", "italy")

	// raw string
	raw := "I am eating $FRUIT and driving ${CAR} in ${COUNTRY}"

	// formatted string
	formatted := os.ExpandEnv(raw)

	// print formatted string
	fmt.Println(formatted)
}
*/

/* =========== 3. Get system and user information =========== */

/* ===== 1. os.Hostname() ===== */
/*
func main() {
	// get `hostname`
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		return
	}
	// print `hostname`
	fmt.Println(hostname, err)
}
*/

/* ===== 2. os.UserHomeDir() ===== */
/*
func main() {
	// get `host` directory
	hostDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}
	// print `host` directory
	fmt.Println(hostDir, err)
}
*/

/* ===== 3. os/user package ===== */
/*
func main() {
	// get current user info
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v\n", currentUser)
	fmt.Printf("err: %v\n", err)
}
*/

/* =========== 4. Working directories =========== */

/* ===== 1. os.Getwd() and os.Chdir ===== */
/*
func main() {
	// get current working directory
	cwd1, err := os.Getwd()
	if err != nil {
		log.Fatal(cwd1)
		return
	}
	fmt.Printf("[before] cwd: %v\n", cwd1)

	// change current working directory
	os.Chdir("..")

	// get current working directory
	cwd2, err := os.Getwd()
	if err != nil {
		log.Fatal(cwd2)
		return
	}
	fmt.Printf("[after] cwd: %v\n", cwd2)

	// change current working directory
	os.Chdir("../..")

	// get current working directory
	cwd3, err := os.Getwd()
	if err != nil {
		log.Fatal(cwd3)
		return
	}
	fmt.Printf("[after again] cwd: %v\n", cwd3)
}
*/

/* ===== 2. os.Executable ===== */
/*
func main() {
	// get executable file path
	exeDir, _ := os.Executable()
	fmt.Println("exeDir", exeDir)
}
*/

/* ===== 3. os.TempDir ===== */
/*
func main() {
	// get temporary directory path
	tempDir := os.TempDir()
	fmt.Println("tempDir", tempDir)
}
*/

/* ===== 4. ioutil.TempDir() and ioutil.TempFile() ===== */

func main() {
	// get temporary directory path
	tempDir := os.TempDir()
	fmt.Println("tempDir", tempDir)

	name, err := ioutil.TempDir(".", "hello")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(name)

	file, err := ioutil.TempFile(".", "index*.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file)
	file.WriteString("<h1>Hello World</h1>")
}
