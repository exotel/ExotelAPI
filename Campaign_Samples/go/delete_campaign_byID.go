package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// Please provide Exotel Sid, API_Key , API Key and API Token from your Exotel account
	exotelSid := "<your Exotel Sid>"
	apiKey := "<your API Key>"
	apiToken := "<your API Token>"

	// Encoding the Exotel Sid API Key and API Token, used in Authorization header
	encoding := b64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiToken))

	url := "https://api.exotel.com/v2/accounts/" + exotelSid + "/campaigns/<campaign SID>"

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("Authorization", "Basic "+encoding)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
