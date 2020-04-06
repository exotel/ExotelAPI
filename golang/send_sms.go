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

	apiKey := "<your-API-key>"
	apiToken := "<your-API-token>"
	exotelSid := "<your-exotel-sid>"
	smsBody := `Hello World!`

	baseURL := fmt.Sprintf("https://%s:%s@api.exotel.in/v1/Accounts/%s/SMS/send", apiKey, apiToken, exotelSid)

	msgData := url.Values{}
	msgData.Set("To", "XXXXXXXXXXX")   //Replace XXXXXXX with To number
	msgData.Set("From", "XXXXXXXXXXX") //Replace XXXXXXX with From number
	msgData.Set("Body", smsBody)

	msgDataReader := strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, baseURL, msgDataReader)
	if err != nil {
		log.Fatal(err)
	}
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
