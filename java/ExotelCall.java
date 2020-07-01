package com.exotel.Connect;

import java.io.IOException;

import javax.naming.ldap.ExtendedResponse;

import com.google.gson.Gson;

import okhttp3.Credentials;
import okhttp3.MultipartBody;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;

class ExotelCall {
  public static String customerNumber = "your-customer-number";
  public static String url = "http://my.exotel.com/";
  public static String accountSid = "your-exotel-account-sid";
  public static String flow_id = "your-flow-id";
  public static String authId = "your-API-Key";
  public static String authToken = "your-API-Token";

  public ExotelResponse connectCustomerToFlow() {
    OkHttpClient client = new OkHttpClient().newBuilder().build();
    RequestBody body = new MultipartBody.Builder().setType(MultipartBody.FORM).addFormDataPart("From", customerNumber)
        .addFormDataPart("Url", url + accountSid + "/exoml/start_voice/" + flow_id).build();

    String credentials = Credentials.basic(authId, authToken);

    Request request = new Request.Builder().url(String.format(ExotelStrings.CONNECT_CUSTOMER_TO_FLOW_URL, accountSid))
        .method("POST", body).addHeader("Authorization", credentials).addHeader("Content-Type", "application/json")
        .build();

    try {
      Response response = client.newCall(request).execute();
      Gson connect = new Gson();
      String res = null;
      try {
        res = response.body().string();
      } catch (IOException e) {
        e.printStackTrace();
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
    } catch (Exception e) {
      e.printStackTrace();
    }

    ExotelFailureResponse cust = new ExotelFailureResponse(0);
    return cust;
  }
}
