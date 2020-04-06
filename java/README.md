EXOTEL-JAVA

# DOCUMENTATION
The documentation for the Exotel API can be found here https://developer.exotel.com/api/#intro.

# SUPPORTED NODEJS VERSIONS
This library supports the following Java implementations:

    Java 8 
    openjdk 8

# INSTALLATION
    1 First, update the apt package index with:
     $sudo apt update
    
    2 Once the package index is updated install the default Java OpenJDK package with:
    $ sudo apt install default-jdk
    
    3 Verify the installation, by running the following command which will print the Java version:
    java -version

# Getting Started
Setup Work
## Put Your Own Credentials Here
    Exotel_sid = 'ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
    API_token = 'yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy'

# Make a Call
```java
class ExotelCall {
	public static String customerNumber = "your-customer-number";
	public static String url = "http://my.exotel.com/";
	public static String exotelsid = "your-exotel-sid";
	public static String flow_id = "your-flow-id";
	public static String apiid = "your-API-Key";
	public static String apitoken = "your-API-token";

	public ExotelResponse connectCustomerToFlow() {
		OkHttpClient client = new OkHttpClient().newBuilder().build();
		RequestBody body = new MultipartBody.Builder().setType(MultipartBody.FORM)
				.addFormDataPart("From", customerNumber)
				.addFormDataPart("Url", url + exotelsid + "/exoml/start_voice/" + flow_id).build();

		String credentials = Credentials.basic(apiid, apitoken);

		Request request = new Request.Builder()
				.url(String.format(ExotelStrings.CONNECT_CUSTOMER_TO_FLOW_URL, exotelsid)).method("POST", body)
				.addHeader("Authorization", credentials).addHeader("Content-Type", "application/json").build();

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

```
# Send an SMS
```java
class ExotelSend {
	public static String customerNumber = "your-customer-number";
	public static String agentNumber    = "your-agent-number";
	public static String url            = "http://my.exotel.com/";
	public static String exotelSid     = "your-exotelsid";
	public static String apiid        = "your-API-Key";
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

```
