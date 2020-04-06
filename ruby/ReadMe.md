
EXOTEL-RUBY

# DOCUMENTATION
The documentation for the Exotel API can be found link:-https://developer.exotel.com/api/#intro

# SUPPORTED RUBY VERSIONS
This library supports the following Ruby implementations:
    Ruby 2.4
    Ruby 2.5
    Ruby 2.6
    Ruby 2.7


# INSTALLATION
To install ruby

    1 First, update the packages index:
    $ sudo apt update
    
    2 Install Ruby by typing:
    $sudo apt install ruby-full
    
    3 To verify that the installation it was successful run the following command which will print the Ruby version:
    $ruby --version

The output will look something like this:
        ruby 2.5.1p57 (2018-03-29 revision 63029) [x86_64-linux-gnu]

# GETTING STARTED
Setup Work:-
# PUT YOUR OWN CREDENTIALS HERE
```ruby
SID   = 'ACxxxxxxxxxxxxxxxxxxxxxxxxxx'
KEY   = 'ACxxxxxxxxxxxxxxxxxxxxxxxxxx'
TOKEN = 'yyyyyyyyyyyyyyyyyyyyyyy'
module Cred

    KEY   = "your-API-key"
    SID   = "your-exotel-sid"
    TOKEN = "your-API-token"
 
    define_singleton_method("connect") {
        URL = "https://#{KEY}:#{TOKEN}@api.exotel.in/v1/Accounts/#{SID}/Calls/connect.json"
       }
    
    define_singleton_method("sms") {
        URL = "https://#{KEY}:#{TOKEN}@api.exotel.in/v1/Accounts/#{SID}/Sms/send.json"
   }
end

```
# MAKE A CALL
```ruby
PUT YOUR APP_ID HERE:-
APP_ID = XXXXXXX
require './config'
require_relative './connect.rb'
include  Cred

Connect.new(Cred.connect, 'To' => 'your-agent-number',  'From' => 'your-customer-number', 'Url' => 'http://my.exotel.com/<exotel-sid>/exoml/start_voice/<app_id>').connect_to_flow
Connect.new(Cred.connect, 'To' => 'your-agent-number', 'From' => 'your-customer-number').connect_to_agent

# SEND AN SMS
require './config'
require_relative './connect.rb'
include  Cred
Connect.new(Cred.sms,'To' => 'your-agent-number', 'From' => 'your-customer-number', 'Body' => 'your-message').connect_to_sms

# HANDLING ERRORS
require 'uri'
require 'net/http'

class NetHTTP

  def self.post(url, params)
    res = Net::HTTP.post_form(URI(url), params)
    puts res.read_body
  rescue TimeoutError => te
    puts te
  rescue StandardError => se
    puts se
  end
end
```

# RESULT
    {"SMSMessage":{"Sid":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "AccountSid":"Exotel","From":"xxxxxxxxx\/EXOTEL",
    "To":"+xxxxxxxxxx",
    "DateCreated":"2020-02-20 15:12:29",
    "DateUpdated":"2020-02-20 15:12:29","DateSent":null,
    "Body":"hello world",
    "Direction":"outbound-api",
    "Uri":"\/v1\/Accounts\/Exotel\/SMS\/Messages\/xxxxxxxxxxxxxxxxxxxxxxxxx.json",
    "ApiVersion":null,"Price":null,
    "Status":"queued",
    "DetailedStatusCode":21010,
    "DetailedStatus":"PENDING_TO_OPERATOR"}}

