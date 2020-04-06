# EXOTEL-PYTHON SAMPLE CODE

# DOCUMENTATION
The documentation for the Exotel API can be found in the below link:
https://developer.exotel.com/api/#intro

# SUPPORTED PYTHON VERSION
This library supports the following Python implementations:


    Python 2.7
    Python 3.4
    Python 3.5
    Python 3.6
    Python 3.7
    Python 3.8

# INSTALLATION
To install Python latest stable version:

01 Start by updating the packages list and installing the prerequisites:
  $ sudo apt update
  $ sudo apt install software-properties-common

02 Next, add the deadsnakes PPA to your sources list:
  $ sudo add-apt-repository ppa:deadsnakes/ppa
When prompted press Enter to continue:
Press [ENTER] to continue or Ctrl-c to cancel adding it.

03 Once the repository is enabled, install Python 3.7 with:
  $sudo apt install python3.7

04 At this point, Python 3.7 is installed on your Ubuntu system and ready to be used. You can verify it by typing: 
  $python3.7 --version
  Python 3.7.3

# GETTING STARTED
# Settings for Exotel API
## Get your SID and token from 
https://my.exotel.com/apisettings/site#api-credentials

    sid         = 'XXXXXX'                                           
    key         = 'XXXXXX'                                           
    token       = 'XXXXXXXXXXXXXXXXXXXXXXXXX'         
    smsurl      = 'https://api.exotel.in/v1/Accounts/<exotel_sid>/Sms/send.json'
    callurl     = 'https://api.exotel.in/v1/Accounts/<exotel_sid>/Calls/connect.json'
    from_no     = 'XXXXXXXXXXX'
    to          = 'XXXXXXXXXXX'
    url         = "http://my.exotel.com/<exotel_sid>/exoml/start_voice/<flow_id>", 

# MAKE A CALL


```python
import requests
from pprint import pprint

from settings import sid, token, callurl,from_no,to


def connect_customer_to_agent(
        sid,
        token,
        agent_no,
        customer_no,
        callerid,
        timelimit=None,
        timeout=None,
        calltype = 'trans'
    ):
    return requests.post(
        callurl,
        auth = (sid, token),
        data = {
            'From': from_no,
            'To'  : to,
          
        }
    )

if __name__ == '__main__':
    r = connect_customer_to_agent(
        sid,
        token,
        agent_no    = "your-agent-number",
        customer_no = "your-customer-number",
        callerid    = "<Your-Exotel-virtual-number>",
        timelimit   = "<time-in-seconds>",  
        timeout     = "<time-in-seconds>",  
        calltype    = "trans"  
    )
    print (r.status_code)
    pprint(r.json())
```
    
# SEND AN SMS
```python
import requests

from pprint import pprint

from settings import sid, token, smsurl,from_no,to

def send_message (sid,token, sms_from, sms_to, sms_body) :
    return requests.post(
        smsurl,
        auth = (sid, token),
        data = {
            'From' : from_no,
            'To'   : to,
            'Body' : "your-message"
        }
    )

if __name__ == '__main__':
    r = send_message(
        sid,
        token,
        sms_from = 'your-agent-number',  
        sms_to   = 'your-customer-number', 
        sms_body = 'your-message'
    )
    print (r.status_code)
    pprint(r.json())
    
```
# Result
HTTO Response : 200

        `{'SMSMessage': {'AccountSid': 'Exotel',
                         'ApiVersion': None,
                         'Body': 'hii',
                         'DateCreated': '2020-02-20 15:54:42',
                         'DateSent': None,
                         'DateUpdated': '2020-02-20 15:54:42',
                         'DetailedStatus': 'PENDING_TO_OPERATOR',
                         'DetailedStatusCode': 21010,
                         'Direction': 'outbound-api',
                         'From': 'xxxxxxxxx/EXOTEL',
                         'Price': None,
                         'Sid': 'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx',
                         'Status': 'queued',
                         'To': '+xxxxxxxxxxxxx',
                         'Uri': '/v1/Accounts/Exotel/SMS/Messages/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.json'}}`


