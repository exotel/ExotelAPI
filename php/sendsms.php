<?php
#Link to developer portal  for send sms https://developer.exotel.com/api/#send-sms 
$post_data = array(
    'From'    => 'xxxxxxxxxx',
    'To'      => 'xxxxxxxxxx',
    'Body'    => 'xxx', 
);
$api_key      = "API_KEY";
$api_token    = "API_TOKEN"; 
$exotel_sid   = "Exotel Sid";
#Replace <subdomain> with the region of your account 
#<subdomain> of Singapore cluster is @api.exotel.com
#<subdomain> of Mumbai cluster is @api.in.exotel.com
$url    = "https://" . $api_key . ":" . $api_token ."@api.exotel.in/v1/Accounts/" . $exotel_sid ."/Sms/send"; 
$ch     = curl_init();
curl_setopt($ch, CURLOPT_VERBOSE, 1);
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
