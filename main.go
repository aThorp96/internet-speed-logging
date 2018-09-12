package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "run.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
}
