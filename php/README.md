# EXOTEL-PHP
# DOCUMENTATION
The documentation for the Exotel API can be found link:https://developer.exotel.com/api/#intro

# SUPPORTED PHP VERSIONS
This library supports the following PHP implementations:
    
    PHP 7.2
    PHP 7.3
    PHP 7.4

# INSTALLATION
    If you are using Apache as your web server to install PHP and Apache PHP module run the following command: $ sudo apt install php libapache2-mod-php 
    
    Once the packages are installed restart the Apache service: $ sudo systemctl restart apache2

# GETTTING STARTED
        put your credentials here
        $api_key = "ACXXXXXXXXXXXXXXXXX" $api_token = "YYYYYYYYYYYYYYYYYY"       	
    	$exotel_sid = "ACXXXXXXXXXXXXXXXXX"

# MAKE A CALL
```php
<?php #Link to developer portal for connect to flow https://developer.exotel.com/api/#call-customer
$post_data = array(
'From' => "your-agent-number", 
'To' => "your-customer-number", 
'Url' => "http://my.exotel.in/exoml/start/<flow_id&gt;", 
'CallType' => "trans" );
$api_key = "your-API-key";
$api_token = "your-API-token";
$exotel_sid = "your-exotel-sid";
#Replace <subdomain> with the region of your account
#<subdomain> of Singapore cluster is @api.exotel.com 
#<subdomain> of Mumbai cluster is @api.in.exotel.com 
$url = "https://" . $api_key . ":" . $api_token . "@<subdomain>/v1/Accounts/" . $exotel_sid . "/Calls/connect"; 
$ch = curl_init(); curl_setopt($ch, CURLOPT_VERBOSE, 1); 
curl_setopt($ch, CURLOPT_URL, $url); 
curl_setopt($ch, CURLOPT_POST, 1); 
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_FAILONERROR, 0); 
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($post_data));
$http_result = curl_exec($ch); 
curl_close($ch);
print "Response = ".print_r($http_result); 
?>

```
# SEND AN SMS
Link to developer portal for send sms https://developer.exotel.com/api/#send-sms

```php
<?php  
$post_data = array( 'From' => 'your-agent-number', 'To' => 'your-customer-number', 'Body' => 'your-message', ); 
$api_key = "your-API-key"; 
$api_token = "your-API-token"; 
$exotel_sid = "your-exotel-sid"; 
#Replace <subdomain> with the region of your account 
#<subdomain> of Singapore cluster is @api.exotel.com 
#<subdomain> of Mumbai cluster is @api.in.exotel.com $url = "https://" . $api_key . ":" . $api_token ."@<subdomain>/v1/Accounts/" . $exotel_sid ."/Sms/send"; $ch = curl_init(); curl_setopt($ch, CURLOPT_VERBOSE, 1); 
curl_setopt($ch, CURLOPT_URL, $url); 
curl_setopt($ch, CURLOPT_POST, 1); 
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1); 
curl_setopt($ch, CURLOPT_FAILONERROR, 0); 
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0); 
curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($post_data)); 
$http_result = curl_exec($ch); curl_close($ch); print "Response = ".print_r($http_result); 
?>

```
# RESULT 
```xml
<?xml version="1.0" encoding="UTF-8"?> 
<SMSMessage> <Sid>xxxxxxxxxxxxxxxxxxxxxxxx</Sid> 
<AccountSid>xxxxxx</AccountSid> 
<From>xxxxxxxxx/EXOTEL</From> 
<To>+xxxxxxxxxxx</To> 
<DateCreated>2020-02-20 18:15:42</DateCreated> 
<DateUpdated>2020-02-20 18:15:42</DateUpdated> 
<DateSent/> <Body>Hii</Body> 
<Direction>outbound-api</Direction>
<Uri>/v1/Accounts/xxxxx/SMS/Messages/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</Uri> 
<ApiVersion/> <Price/> <Status>queued</Status> 
<DetailedStatusCode>21010</DetailedStatusCode> 
<DetailedStatus>PENDING_TO_OPERATOR</DetailedStatus> 
</SMSMessage> 

```
