<?php 
#Link to developer portal  for connect to agent https://developer.exotel.com/api/#call-agent
$post_data = array(
    'From'     => "your-agent-number",
    'To'       => "your-customer-number",
    'CallType' => "promo" 
); 
$api_key     = "your-API key"; 
$api_token   = "your-API token"; 
$exotel_sid  = "your-exotel-sid";
#Replace <subdomain> with the region of your account
#<subdomain> of Singapore cluster is @api.exotel.com
#<subdomain> of Mumbai cluster is @api.in.exotel.com 
$url = "https://" . $api_key .  ":"  . $api_token . "@<subdomain>/v1/Accounts/" . $exotel_sid . "/Calls/connect"; 
$ch  = curl_init();
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

