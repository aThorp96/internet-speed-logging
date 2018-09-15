package main

import (
	"log"
	"os/exec"
)

func main() {

	// Run speed test
	cmd := exec.Command("/usr/bin/speedtest", "--json")
	// Capture ouptut
	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("speedtest not able to run. Make sure it is installed. %v", err)
	}

	speedTest := initSpeed(output)

	MakeRequest(speedTest)

}
