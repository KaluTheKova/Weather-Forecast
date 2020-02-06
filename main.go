package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Save this so it can be used again
var myClient = &http.Client{Timeout: 10 * time.Second}

// APIKey that required registering. Horrible.
var YourAPIKey = "72897d5c8e06a55ebdf99d0d52ce8bc8"

// Struct for weatherdata
type weatherAPI struct {
	Coord struct {
		Longitude float64 `json:"lon"`
		Latitude  float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

// Custom Stringer for weatherAPI output
func (w weatherAPI) String() string {
	return fmt.Sprintf("Coordinates: {Longitude: %v, Latitude: %v}\nWeather is: %v", w.Coord.Longitude, w.Coord.Latitude, w.Weather[0])
}

// Get JSON data from url, return results in weatherAPI struct
func GetWeatherJSON(url string) (*weatherAPI, error) {

	response, err := myClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Get JSON from url failed: %v", err)
	}
	// Close this
	defer response.Body.Close()

	// Decodes response to target structure. PROBLEM: JSON may be in faulty format and decoder does not work.
	var data weatherAPI
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Read city from user
func ReadCity() string {
	fmt.Println("Enter city name. Remember to start with Capital letter.")
	var input string

	fmt.Scanln(&input) // Does not work with whitespaces. Horrible.
	fmt.Println("Your input was:", input)
	return input
}

func main() {

	// Ask user for city
	city := ReadCity()

	// Add cityname and APIKey into url
	url := "http://api.openweathermap.org/data/2.5/weather?APPID=" + YourAPIKey + "&q=" + city
	fmt.Println(url)

	// Get data
	data, err := GetWeatherJSON(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
