package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println("Executing Processes")

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		log.Fatal(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
