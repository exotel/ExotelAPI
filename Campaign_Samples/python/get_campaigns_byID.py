import requests, base64, json

accountSid = "<your Exotel SID>"
accountToken = "<your Exotel token>"
encoding = base64.b64encode(accountSid + ":" + accountToken)
url = "https://api.exotel.com/v2/accounts/"+accountSid+"/campaigns/<campaign sid>"

payload = ""
headers = {
    'Authorization': "Basic " + encoding
    }

response = requests.request("GET", url, data=payload, headers=headers)

print(response.text)
print(json.dumps(json.loads(response.text), indent = 4, sort_keys = True))

