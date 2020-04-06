key="API Key"
sid="Exotel Sid"
token="API Token"
from="xxxxxxxxxxxxxx"
to="xxxxxxxxxx"


const r = x =>Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '')

const axios = require('axios')

url="https://"+key+":"+token+"@api.exotel.in/v1/Accounts/"+sid+"/Calls/connect.json"
axios.post(url, 
   r({
  "From": from,
  "To": to,
  "CallerId":'xxxxxxxxx',
  "CallerType": 'trans',
   "Url" :'http://my.exotel.com/Exotel/exoml/start_voice/26555',

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