<?php

$curl = curl_init();
$accountSid = "<your Exotel Sid>";
$apiKey = "<your API Key>";
$apiToken = "<your API Token>";
$encoding = base64_encode($apiKey .":". $apiToken);
curl_setopt_array($curl, array(
  CURLOPT_URL => "https://api.exotel.com/v2/accounts/".$accountSid."/campaigns",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 30,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "POST",
  CURLOPT_POSTFIELDS => '{
                "campaigns": [{
                "caller_id": "0XXXXXXXXX1",
                "url": "http://my.exotel.in/exoml/start/189810",
                "from": ["+91XXXXXXXXX4", "+91XXXXXXXXX5"],
                "schedule": {"send_at": "2018-11-08T14:20:00+05:30"},
                "status_callback": "http://<callback custom domain>/1gvta9f1",
                "call_status_callback": "http://<callback custom domain>/1gvta9f1"}]}',
  CURLOPT_HTTPHEADER => array(
    "Authorization: Basic ".$encoding,
    "Content-Type: application/json"
  ),
));

$response = curl_exec($curl);
$err = curl_error($curl);

curl_close($curl);

if ($err) {
  echo "cURL Error #:" . $err;
} else {
  echo $response;
}
