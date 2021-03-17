package main

import (
	"fmt"
	"os"
	"os/exec"
)

func isInstalled(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v " + name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func subExecute(program string, args ...string) ([]byte, error) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return cmd.CombinedOutput()
}
