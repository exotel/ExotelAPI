package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	// Please provide ExotelSid, APIToken from your Exotel account
	exotelSid := "<your Exotel Sid>"
	apiKey := "<your API Key>"
	apiToken := "<your API Token>"

	// Encoding the Exotel Sid ,API_Key , API Token, used in Authorization header
	encoding := b64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiToken))

	url := "https://api.exotel.com/v2/accounts/" + exotelSid + "/campaigns"

	payload := strings.NewReader(`{ 
		"campaigns": [{ 
			"caller_id": "0XXXXXXXXX1", 
			"url": "http://my.exotel.com/exoml/start_voice/208287", 
			"from": [ "+91XXXXXXXXX4", "+91XXXXXXXXX5" ], 
			"status_callback": "http://<callback custom domain>/1gvta9f1", 
			"call_status_callback": "http://<callback custom domain>/1gvta9f1"
			}]
		}`)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+encoding)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
