EXOTEL-NODEJS


# DOCUMENTATION
The documentation for the Exotel API can be found here https://developer.exotel.com/api/#intro.


# SUPPORTED NODEJS VERSIONS
This library supports the following Nodejs implementations:
    
    Nodejs v8

# INSTALLATION
    To get this version, you can use the apt package manager. Refresh your local package index by typing:
    $ sudo apt update
    Install Node.js from the repositories:
    $ sudo apt install nodejs
    If the package in the repositories suits your needs, this is all you need to do to get set up with Node.js. In most cases, you’ll also want to also install npm, the Node.js package manager. You can do this by typing:
    $ sudo apt install npm
    This will allow you to install modules and packages to use with Node.js.
    Because of a conflict with another package, the executable from the Ubuntu repositories is called nodejs instead of node. Keep this in mind as you are running software.
    
    To check which version of Node.js you have installed after these initial steps, type:
    
    nodejs -v

# Getting Started
Setup Work
## Put Your Own Credentials here
    Exotel_sid = 'ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
    API_Token = 'yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy'

# Make a Call

 ```javascript
   key="xxxxxxxx"
    sid="xxxxxxxx"
    token="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    from="xxxxxxxxxxx"
    to="xxxxxxxxxxxx"
```

```javascript
const formUrlEncoded = x =>Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '')
const axios = require('axios')
url="https://"+key+":"+token+"@api.exotel.in/v1/Accounts/"+sid+"/Calls/connect"
axios.post(url, 
   formUrlEncoded({
  "From": from,
  "To": to,
  "CallerId":'xxxxxxxxx',
  "CallerType": 'promo',
}),
{   
    withCredentials: true,
    headers: {
        "Accept":"application/x-www-form-urlencoded",
        "Content-Type": "application/x-www-form-urlencoded"
    }
  },
  )
.then((res) => {
  console.log(`statusCode: ${res.statusCode}`)
  console.log(res)
})
.catch((error) => {
  console.error(error)
})
```

# Send an SMS
```javascript
key="xxxxxxx"
sid="xxxxxxx"
token="xxxxxxxxxxxxxxxxxxxxx"
from="xxxxxxxxx"
to="xxxxxxxxxxx"
body="Good Evening"

const formUrlEncoded = x =>Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '')

const axios = require('axios')
url="https://"+key+":"+token+"@api.exotel.in/v1/Accounts/"+sid+"/Sms/send.json"
axios.post(url, 
   formUrlEncoded({
  "From": from,
  "To": to,
  "Body":body
}),
{   
    withCredentials: true,
    headers: {
        "Accept":"application/x-www-form-urlencoded",
        "Content-Type": "application/x-www-form-urlencoded"
    }
  },
  )
.then((res) => {
  console.log(`statusCode: ${res.statusCode}`)
  console.log(res)
})
.catch((error) => {
  console.error(error)
})

```

# RESULT
HTTO Response Code: 200 
   ```xml
 <?xml version="1.0" encoding="UTF-8"?>\n
    <TwilioResponse>\n <Call>\n  
    <Sid>xxxxxxxxxxxxxxxxxxxxxxxxxxxxx</Sid>\n  
    <ParentCallSid/>\n  
    <DateCreated>2020-02-24 13:21:32</DateCreated>\n  
    <DateUpdated>2020-02-24 13:21:32</DateUpdated>\n  
    <AccountSid>xxxxxxx</AccountSid>\n  
    <To>+xxxxxxxxxx</To>\n  
    <From>+xxxxxxxxxx</From>\n  
    <PhoneNumberSid>+xxxxxxxxxxx</PhoneNumberSid>\n  
    <Status>in-progress</Status>\n  
    <StartTime>2020-02-24 13:21:32</StartTime>\n  
    <EndTime/>\n  <Duration/>\n  <Price/>\n  
    <Direction>outbound-api</Direction>\n  
    <AnsweredBy/>\n  
    <ForwardedFrom/>\n  
    <CallerName/>\n  
    <Uri>/v1/Accounts/xxxxxx/Calls/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</Uri>\n  
    <RecordingUrl/>\n 
    <SubResourceUris>\n   
    <Legs>/v1/Accounts/xxxxxxx/Calls/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/Legs</Legs>\n  </SubResourceUris>\n </Call>\n</TwilioResponse>\n'
```


  
