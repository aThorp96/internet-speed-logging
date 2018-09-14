package main

import (
	"fmt"
	"os/exec"
	//"time"
)

func main() {

	// Run speed test
	cmd := exec.Command("/usr/bin/speedtest", "--json")
	// Capture ouptut
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	speedTest := initSpeed(output)

	MakeRequest(speedTest)

}
