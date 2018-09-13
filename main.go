package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	//"time"
)

type SpeedTest struct {
	Down   float32 `json:"download"`
	Up     float32 `json:"upload"`
	Server Server  `json:"server"`
	Time   string  `json:"timestamp"`
}

type Server struct {
	Name string `json:"name"`
}

func main() {

	cmd := exec.Command("/bin/bash", "run.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	//statePath := "state.json"
	speedPath := "speed.json"
	jsonFile, err := os.Open(speedPath)
	if err != nil {
		fmt.Println(err)
	}

	var speedTest SpeedTest

	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &speedTest)

	fmt.Printf("Upload speed   : %f\n", speedTest.Up)
	fmt.Printf("Download speed : %f\n", speedTest.Down)
	fmt.Println("Server        : " + speedTest.Server.Name)
	fmt.Println("Timestamp     : " + speedTest.Time)

	defer jsonFile.Close()
}
