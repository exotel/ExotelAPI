package com.exotel.Connect;

import java.io.IOException;

import javax.naming.ldap.ExtendedResponse;

import com.google.gson.Gson;

import okhttp3.Credentials;
import okhttp3.MultipartBody;
import okhttp3.MultipartBody.Builder;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;

enum SmsType {
	TRANSACTIONAL, PROMOTIONAL;
}

class ExotelSend {
	public static String customerNumber = "your-customer-number";
	public static String agentNumber    = "your-agent-number";
	public static String url            = "http://my.exotel.com/";
	public static String exotelSid     = "your-exotel-sid";
	public static String apiid         = "your-API-Key";
	public static String apitoken      = "your-API-token";
	public static String msg            = "your-message";


	public ExotelResponse SendUnicode() {
		OkHttpClient client = new OkHttpClient().newBuilder().build();
		RequestBody body = new MultipartBody.Builder().setType(MultipartBody.FORM)
				.addFormDataPart("From", customerNumber)
				.addFormDataPart("To",agentNumber )
				.addFormDataPart("SmsType", toString())
				.addFormDataPart("EncodingType", "plain")
				.addFormDataPart("Body",body);

		String credentials = Credentials.basic(apiid, apitoken);

		Request request = new Request.Builder()
				.url(String.format(ExotelStrings.SEND_UNICODE_URL, exotelSid)).method("POST", body)
				.addHeader("Authorization", credentials).addHeader("Content-Type", "application/json").build();

		try {
			Response response = client.newCall(request).execute();
			Gson connect = new Gson();
			String res = null;
			try {
				res = response.body().string();
			} catch (IOException error) {
				error.printStackTrace();
			}

			ExotelResponse SuccessResponse = connect.fromJson(res, ExotelResponse.class);

			int status = response.code();

			if (status == 200) {
				ExotelSuccessResponse cust = new ExotelSuccessResponse(0);
				return cust;
			} else {
				ExotelFailureResponse cust = new ExotelFailureResponse(0);
				return cust;
			}
		} catch (Exception error) {
			error.printStackTrace();
		}

		ExotelFailureResponse cust = new ExotelFailureResponse(0);
		return cust;
	}
}