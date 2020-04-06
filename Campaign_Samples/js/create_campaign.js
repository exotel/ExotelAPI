var request = require("request");
var accountSid = "<your Exotel Sid>";
var apiKey = "<Your API Key>"
var apiToken = "<Your API Token>";
var encoding = Buffer.from(apiKey + ':' + apiToken).toString('base64')
var options = { method: 'POST',
  url: 'https://api.exotel.com/v2/accounts/'+ accountSid + '/campaigns',
  headers: 
   { 
     Authorization: 'Basic ' + encoding,
     'Content-Type': 'application/json' },
  body: 
   { campaigns: 
      [ { caller_id: '0XXXXXXXXX1',
          url: 'http://my.exotel.in/start/189810',
          from: [ '+91XXXXXXXXX4', '+91XXXXXXXXX5' ],
          schedule: { send_at: '2018-11-08T14:20:00+05:30' },
          status_callback: 'http://<callback custom domain>/1gvta9f1',
          call_status_callback: 'http://<callback custom domain>/1gvta9f1' } ] },
  json: true };

request(options, function (error, response, body) {

  if (error) throw new Error(error);
  console.log(body.error_data)
  console.log(body)
});
