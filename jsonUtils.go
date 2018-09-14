package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	//"os"
)

// For unpacking speedtest json
type SpeedTest struct {
	Down   float64 `json:"download"`
	Up     float64 `json:"upload"`
	Server Server  `json:"server"`
	Time   string  `json:"timestamp"`
}

type Server struct {
	Name string `json:"name"`
}

// Initializes the speedTest struct
func initSpeed(stdout io.ReadCloser) SpeedTest {

	var speedTest SpeedTest

	bytes, _ := ioutil.ReadAll(stdout)
	json.Unmarshal(bytes, &speedTest)

	//Verify input
	fmt.Printf("Upload speed   : %f\n", speedTest.Up)
	fmt.Printf("Download speed : %f\n", speedTest.Down)
	fmt.Println("Server         : " + speedTest.Server.Name)
	fmt.Println("Timestamp      : " + speedTest.Time)

	return speedTest
}

type State struct {
}
