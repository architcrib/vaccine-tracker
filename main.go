package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type url string

func main() {
	pinCode := 201012 // os.Args[0]
	// districtID := 650 // Noida, UP
	districtID := 651 // Ghaziabad, UP
	urlByPin := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=%v&date=%v"
	urlByDistrict := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=%v&date=%v"
	// Build the URLs for the month
	urlsForPinCode := []string{
		fmt.Sprintf(urlByPin, pinCode, "01-05-2021"),
		fmt.Sprintf(urlByPin, pinCode, "07-05-2021"),
		fmt.Sprintf(urlByPin, pinCode, "13-05-2021"),
		fmt.Sprintf(urlByDistrict, districtID, "01-05-2021"),
		fmt.Sprintf(urlByDistrict, districtID, "07-05-2021"),
		fmt.Sprintf(urlByDistrict, districtID, "13-05-2021"),
	}
	c := make(chan []ValidCenter)
	for _, url := range urlsForPinCode {
		go FindValidCenter(url, c)
	}
	for range urlsForPinCode {
		validCenterElement := <-c
		// Re format this value into a clear and comprehensible message string that has to be
		// passed to my Email.
		for _, validCenter := range validCenterElement {
			fmt.Println(validCenter.CenterName, " - ", validCenter.Vaccine, " - ", validCenter.AvailableCapacity, " - ", validCenter.Date)
		}
	}
}
func FindValidCenter(url string, c chan []ValidCenter) {
	cowinResponse := CowinResponse{}
	validCenters := []ValidCenter{}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while calling the coWIN API for the results")
		c <- validCenters
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		json.Unmarshal([]byte(bodyString), &cowinResponse)
	}
	cowinResponse.FilterValidCenter(&validCenters)
	c <- validCenters
}
