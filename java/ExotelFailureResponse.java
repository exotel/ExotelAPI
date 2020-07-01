package com.exotel.Connect;

class ExotelFailureResponse extends ExotelResponse {
  public ExotelFailureResponse Calls;
  public int httpStatus;
  public String badrequest;
  public String actionforbidden;
  public String invalidfromspecified;

  public ExotelFailureResponse(int httpStatus) {
    this.httpStatus = httpStatus;
  }
}
