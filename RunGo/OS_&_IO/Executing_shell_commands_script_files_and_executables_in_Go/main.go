package main

import (
	"fmt"
	"os/exec"
)

/*
=============
ALL FUNCTIONS

1). exec.LookPath(pathname/or cmd name string)
2). exec.Cmd{}
3). exec.Command(name string,args ...string)
4). exec.Start()
5). exec.Wait()
6). exec.Run()
7). cmd.Output()
8). cmd.CombinedOutput()

=============
*/

/* ============== 1) LookPath function ============== */
/*
func main() {
	// find `go` executable path
	printPath("git")
	printPath("hello")
	printPath("go")
}
func printPath(cmd string) {
	goExecPath, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(os.Executable())
		fmt.Println("Go Executable:", goExecPath)
	}
} */

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
	fmt.Println()
	fmt.Println("Args:-", cmdGoVer.Args)
	fmt.Println("Dir:-", cmdGoVer.Dir)
	fmt.Println("Env:-", cmdGoVer.Env)
	fmt.Println("ExtraFiles:-", cmdGoVer.ExtraFiles)
	fmt.Println("Path:-", cmdGoVer.Path)
	fmt.Println("Process:-", cmdGoVer.Process)
	fmt.Println("ProcessState:-", cmdGoVer.ProcessState)
	fmt.Println("Stderr:-", cmdGoVer.Stderr)
	fmt.Println("Stdin:-", cmdGoVer.Stdin)
	fmt.Println("Stdout:-", cmdGoVer.Stdout)
	fmt.Println(cmdGoVer.SysProcAttr)
	fmt.Println()
	// see command represented by `cmdGoVer`
	fmt.Println(cmdGoVer.String())

	// run `go version` command
	if err := cmdGoVer.Run(); err != nil {
		log.Fatal(err)
	}
} */

/* ============== 4) Start and Wait methods ============== */
/*
func main() {
	fmt.Println(exec.LookPath("./sleep.sh"))
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
*/

/* ============== executable/main.exe  ============== */

/*
func main() {
	fmt.Println("`HELLO` Env:-", os.Getenv("HELLO"))
	for index, arg := range os.Args {
		fmt.Printf("Arg %v). %v\n", index, arg)
	}
}
*/
/*
func main() {
	mainExectablePath, err := exec.LookPath("./executable/main.exe")
	if err != nil {
		log.Fatal("mainExecutablePath:-", err)
	}
	stdOut, err := os.OpenFile("./std-streams/out.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("stdOut:-", err)
	}
	stdErr, err := os.OpenFile("./std-streams/err.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("stdErr:-", err)
	}

	cmd := &exec.Cmd{
		Path:   mainExectablePath,
		Args:   []string{mainExectablePath, "First Arg", "Second Arg"},
		Env:    []string{"HELLO=PRINCE"},
		Stdout: stdOut,
		Stderr: stdErr,
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
} */

/* func main() {
	p, _ := exec.LookPath("cargo")
	fmt.Println(p)
	cmd := &exec.Cmd{
		Path:   p,
		Args:   []string{p, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
*/

/* ============== Command Function  ============== */
/*
func main() {
	// construct `go version` command
	cmd := exec.Command("go", "version")

	// configure `Stdout` and `Stderr`
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
*/

/* ============== Output method  ============== */
/*
func main() {
	// construct `go version` command
	cmd := exec.Command("go", "version")

	// run command
	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Output: %s\n", output)
	}

	cmd1 := exec.Command("go", "version1")

	// run command
	output, err := cmd1.Output()
	fmt.Printf("Output: %s\n", output)
	fmt.Println("Error:", err)
} */

/* ============== CombinedOutput method  ============== */

func main() {
	// construct `go version` command
	cmd := exec.Command("go", "version")

	// run command
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Output: %s\n", output)
	}

	cmd1 := exec.Command("go", "version1")

	// run command
	output, err := cmd1.CombinedOutput()
	fmt.Printf("Output: %s\n", output)
	fmt.Println("Error:", err)
}
