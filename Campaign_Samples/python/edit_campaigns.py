import requests, base64, json

accountSid = "<your Exotel SID>"
apiKey = "<your API key>"
apiToken = "<your API token>"
encoding = base64.b64encode(apiKey + ":" + apiToken)
url = "https://api.exotel.com/v2/accounts/"+accountSid+"/campaigns/<campaign SID>"

payload = json.dumps({ "campaigns": [{ 
					"caller_id": "0XXXXXXXXX1", 
					"url": "http://my.exotel.com/exoml/start_voice/208287", 
					"from": [ "+91XXXXXXXXX4", "+91XXXXXXXXX5" ], 
					"status_callback": "http://<callback custom domain>/1gvta9f1", 
					"call_status_callback": "http://<callback custom domain>/1gvta9f1"}]
				})
headers = {
    'Content-Type': "application/json",
    'Authorization': "Basic " + encoding
    }

response = requests.request("PUT", url, data=payload, headers=headers)

print(response.text)
print(json.dumps(json.loads(response.text), indent = 4, sort_keys = True))
