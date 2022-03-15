package main

import (
	"fmt"
	"os"
	"os/exec"
)

/* ============== 1) LookPath function ============== */
/*
func main() {
	// find `go` executable path
	goExecPath, err := exec.LookPath("./hello")
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Println(os.Executable())
		fmt.Println("Go Executable:", goExecPath)
	}
}
*/

/* ============== 2) Cmd Structure ============== */
/*
func main() {
	// get `go` executable path
	goExecutable, _ := exec.LookPath("go")

	// `go version` command
	cmdGoVer := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// see command represented by `cmdGoVer`
	fmt.Println(cmdGoVer.String())
}
*/

/* ============== 3) Run method ============== */
/*
func main() {
	// get `go` executable path
	goExecutable, _ := exec.LookPath("go")

	// `go version` command
	cmdGoVer := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// see command represented by `cmdGoVer`
	fmt.Println(cmdGoVer.String())

	// run `go version` command
	if err := cmdGoVer.Run(); err != nil {
		log.Fatal(err)
	}
}
*/

/* ============== 4) Start and Wait methods ============== */

func main() {
	// construct `sleep.sh` command
	cmd := &exec.Cmd{
		Path:   "./sleep.sh",
		Args:   []string{"./sleep.sh", "3"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// run `cmd` in background
	err := cmd.Start()
	fmt.Println(err)

	// do something else
	// for i := 1; i < 300000; i++ {
	// 	fmt.Println(i)
	// }

	// wait `cmd` until it finishes
	err = cmd.Wait()
	fmt.Println(err)
}
