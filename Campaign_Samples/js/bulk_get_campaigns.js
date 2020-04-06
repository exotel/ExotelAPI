var request = require("request");
var accountSid = "<your Exotel Sid>";
var apiKey = "<Your API Key>"
var apiToken = "<Your API Token>";
var encoding = Buffer.from(apiKey + ':' + apiToken).toString('base64')
var options = { method: 'GET',
  url: 'https://api.exotel.com/v2/accounts/'+ accountSid +'/campaigns',
  qs: { page_size: '1', page: '1' },
  headers: 
   { 
     Authorization: 'Basic ' + encoding,
     'Content-Type': 'application/json' } };

request(options, function (error, response, body) {
  if (error) throw new Error(error);

  console.log(body);
});
