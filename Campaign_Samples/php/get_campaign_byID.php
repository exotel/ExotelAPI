<?php

$curl = curl_init();
$accountSid = "<your Exotel Sid>";
$apiKey = "<your API Key>";
$apiToken = "<your API Token>";
$encoding = base64_encode($apiKey .":". $apiToken);
curl_setopt_array($curl, array(
  CURLOPT_URL => "https://api.exotel.com/v2/accounts/".$accountSid."/campaigns/<campaign SID>",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 30,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "GET",
  CURLOPT_POSTFIELDS => "",
  CURLOPT_HTTPHEADER => array(
    "Authorization: Basic ".$encoding
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
