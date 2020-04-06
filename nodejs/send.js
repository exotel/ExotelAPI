key="API Key"
sid="Exotel Sid"
token="API Token"
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