package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"strings"
)

func main() {

	apiToken := "API_Token"
	exotelSid := "Exotel Sid"
	apiKey := "API_Key"

	baseURL := fmt.Sprintf("https://%s:%s@api.exotel.com/v1/Accounts/%s/Calls/connect.json", apiKey, apiToken, exotelSid)

	callData := url.Values{}
	callData.Set("To", "xxxxxxxxx")   //Replace XXXXXXXX with To number.
	callData.Set("From", "xxxxxxxxxx") //Replace XXXXXXX with From number.
	callData.Set("CallerType", "trans")

	callDataReader := strings.NewReader(callData.Encode())
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, baseURL, callDataReader)
	req.SetBasicAuth(exotelSid, apiToken)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
