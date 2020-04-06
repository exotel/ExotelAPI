var request = require("request");
var accountSid = "<your Exotel Sid>";
var apiKey = "<Your API Key>"
var apiToken = "<Your API Token>";
var encoding = Buffer.from(apiKey + ':' + apiToken).toString('base64')

var options = { method: 'DELETE',
  url: 'https://api.exotel.com/v2/accounts/'+ accountSid +'/campaigns/<campaign SID>',
  headers: 
   { 
     Authorization: 'Basic ' + encoding 
    } };

request(options, function (error, response, body) {
  if (error) throw new Error(error);

  console.log(body);
  
});
