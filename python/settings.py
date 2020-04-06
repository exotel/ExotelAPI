# Settings for Exotel API

# Get your SID and token from https://my.exotel.com/apisettings/site#api-credentials

'''
sid     - this represnet your exotel sid
key     - this represent your exotel key
token   - this represent your exotel token
from_no - agent number    (the person who is making a call)
to      - customer number (the person to whom the call is made)

'''
sid         = ' Exotel Sid'                                           
key         = 'API Key'                                           
token       = 'API Token'         
smsurl      = 'https://api.exotel.in/v1/Accounts/<exotel_sid>/Sms/send.json'
callurl     = 'https://api.exotel.in/v1/Accounts/<exotel_sid>/Calls/connect.json'
from_no     = 'XXXXXXXXXXX'
to          = 'XXXXXXXXXXX'
url         = "http://my.exotel.com/<exotel_sid>/exoml/start_voice/<flow_id>", 