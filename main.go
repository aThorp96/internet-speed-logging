package main

import (
	"fmt"
	"os/exec"
	//"time"
)

func main() {

	// Run speed test
	cmd := exec.Command("/bin/bash", "run.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	speedTest := initSpeed()
	MakeRequest(speedTest)

}
