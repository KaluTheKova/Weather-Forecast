package weathergetter

import (
	"encoding/json"
	"fmt"
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

// Get JSON data from url, return results in weatherAPI struct
func GetWeatherJSON(url string) (*weatherAPI, error) {

	response, err := myClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Get JSON from url failed: %v", err)
	}
	// Close this
	defer response.Body.Close()

	/*
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(body))
	*/

	// Decodes response to target structure. PROBLEM: JSON may be in faulty format and decoder does not work.
	var data weatherAPI
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
