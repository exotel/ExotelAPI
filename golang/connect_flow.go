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

	baseURL := fmt.Sprintf("https://%s:%s@api.exotel.com/v1/Accounts/%s/Calls/connect.json", apiKey, apiToken, exotelSid)

	callData := url.Values{}
	callData.Set("From", "XXXXXXXXXX") //Replace XXXXXXXXX with From number.
	callData.Set("Url", "http://my.exotel.com/<exotel_sid>/exoml/start_voice/<flow_id>")
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
