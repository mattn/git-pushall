package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := append([]string{"push"}, os.Args[1:]...)
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
		os.Exit(1)
	}
	os.Exit(0)
}
