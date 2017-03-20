package main

import (
	"fmt"
	"github.com/mcronce/sevrest"
	"strconv"
)

func main() {
	// Create Client and Login
	var c = sevrest.Client("http://localhost:8080/api/v1")
	var err = c.Auth("admin", "yourpassword")
	if err != nil {
		fmt.Printf(err.Error())
	}

	// Data for creating a new device
	device := map[string]string{
		"name":             "Test Device O",
		"description":      "Test Description 1",
		"ipAddress":        "127.0.0.1",
		"pollingFrequency": "300",
	}

	// We will get a full device JSON back, but we only care to parse the deviceId in the response
	type CreateDeviceResponse struct {
		DeviceId int `json:"id"`
	}
	var respDevice CreateDeviceResponse

	// Create the device
	resp, err := c.Post("devices", device)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
	fmt.Println("StatusCode: ", resp.StatusCode)
	err = resp.Decode(&respDevice)
	sevrest.PrettyPrint(respDevice) // We will only see the id Field

	// Otherwise we'll just a map for our response
	var respMap map[string]interface{}

	// Get our newly created device
	resp, err = c.Get("devices/" + strconv.Itoa(respDevice.DeviceId))
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
	fmt.Println("StatusCode: ", resp.StatusCode)
	err = resp.Decode(&respMap)
	sevrest.PrettyPrint(respMap)

	// Get all devices
	resp, err = c.Get("/devices")
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
	fmt.Println("StatusCode: ", resp.StatusCode)
	err = resp.Decode(&respMap)
	sevrest.PrettyPrint(respMap)
}