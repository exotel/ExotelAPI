package com.exotel.Connect;

import okhttp3.Credentials;
import okhttp3.MultipartBody;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;

import java.io.IOException;

import com.google.gson.Gson;

public class SendUnicode {
  public static void main(String[] args) {
    ExotelSend send = new ExotelSend();
    ExotelResponse res = send.SendUnicode();
  }
}
