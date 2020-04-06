using System;
using System.Collections.Generic;
using System.IO;
using System.Net;



namespace ExotelSDK
{
	public class Connect
	{
		private string SID = null;
		private string KEY = null;
		private string TOKEN = null;

		public Connect (string SID,string key, string token)
		{
			this.SID   = "your-exotel-sid";
			this.KEY   = "your-API-key";
			this.TOKEN = "your-API-token";
		}
		public string connectCustomerToApp (string from, string url, string callerID,
		                                     string callType, string timeLimit = null,
		                                     string timeOut = null, string statusCallback = null,
		                                     string customfield = null)
		{
			Dictionary<string, string> postValues = new Dictionary<string, string> ();
			postValues.Add ("From", "09893638575");
			postValues.Add ("Url","http://my.exotel.com/Exotel/exoml/start_voice/261918");
			//postValues.Add ("CallerID", "xxxxxxxxxxx");
			postValues.Add ("CallType", "trans");
			if (timeLimit != null) {
				postValues.Add ("TimeLimit", timeLimit);
			}
			if (timeOut != null) {
				postValues.Add ("TimeOut", timeOut);
			}

			if (statusCallback != null) {
				postValues.Add ("StatusCallback", statusCallback);
			}
			
			
			String postString = "";

			foreach (KeyValuePair<string, string> postValue in postValues) {
				postString += postValue.Key + "=" + WebUtility.UrlEncode (postValue.Value) + "&";
			}
			postString = postString.TrimEnd ('&');

			return(this.sendRequest (postString));
		}

		private string sendRequest (string postString)
		{
			
			 ServicePointManager.ServerCertificateValidationCallback = delegate {
				return true;
			};
			string smsURL = "https://api.exotel.in/v1/Accounts/Exotel/Calls/connect";
			HttpWebRequest objRequest = (HttpWebRequest)WebRequest.Create (smsURL);
			objRequest.Credentials = new NetworkCredential (this.KEY, this.TOKEN);
			objRequest.Method = "POST";
			objRequest.ContentLength = postString.Length;
			objRequest.ContentType = "application/x-www-form-urlencoded";
			                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               
			StreamWriter opWriter = null;
			opWriter = new StreamWriter (objRequest.GetRequestStream ());
			opWriter.Write (postString);
			opWriter.Close ();

			                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            
			HttpWebResponse objResponse = (HttpWebResponse)objRequest.GetResponse ();
			string postResponse = null;
			using (StreamReader responseStream = new StreamReader(objResponse.GetResponseStream())) {
			}

			return (postResponse);
		}
		public static void Main (string[] args)
		{
			Connect c = new Connect ("exotel_sid","API_Key", "API_Token");
	    	string response = c.connectCustomerToApp("To", "http://my.exotel.in/exoml/start/261918","From","trans");
		Console.WriteLine(response);
		}
	}
}
