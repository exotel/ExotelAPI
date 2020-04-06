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

