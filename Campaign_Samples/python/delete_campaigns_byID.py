import requests, base64, json
accountSid = "<your Exotel SID>"
apiKey = "<your API key>"
apiToken = "<your API token>"
encoding = base64.b64encode(apiKey + ":" + apiToken)

url = "https://api.exotel.com/v2/accounts/"+ accountSid +"/campaigns/<campaign SID>"

payload = ""
headers = {
    'Authorization': "Basic " + encoding
    }

response = requests.request("DELETE", url, data=payload, headers=headers)

print(response.text)
print(json.dumps(json.loads(response.text), indent = 4, sort_keys = True))
