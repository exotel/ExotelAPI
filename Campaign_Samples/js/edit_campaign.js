var request = require("request");
var accountSid = "<your Exotel Sid>";
var apiKey = "<Your API Key>"
var apiToken = "<Your API Token>";
var encoding = Buffer.from(apiKey + ':' + apiToken).toString('base64')

var options = { method: 'PUT',
  url: 'https://api.exotel.com/v2/accounts/'+ accountSid +'/campaigns/<campaign SID>',
  headers: 
   {
     Authorization: 'Basic ' + encoding,
     'Content-Type': 'application/json' 
    },
  body: 
   { campaigns: 
      [ { caller_id: '0XXXXXXXXX1',
          url: 'http://my.exotel.in/exoml/start/189810',
          from: [ '+91XXXXXXXX4', '+91XXXXXXXXX5' ],
          schedule: { send_at: '2018-11-09T10:00:00+05:30' },
          status_callback: 'http://<callback custom domain>/1gvta9f1',
          call_status_callback: 'http://<callback custom domain>/1gvta9f1' } ] },
  json: true };

request(options, function (error, response, body) {
  if (error) throw new Error(error);

  console.log(body);
  console.log(JSON.stringify(body))
});
