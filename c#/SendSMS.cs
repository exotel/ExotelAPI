using System;
using System.Collections.Generic;
using System.IO;

using System.Net;


namespace ExotelSDK
{
	public class SendSMS
	{
		private string SID = null;
		private string key = null;
		private string token =null ;

		public SendSMS (string SID,string key,string token)
		{
			this.SID = "your-exotel-sid";
			this.key = "your-exotel-key";
			this.token = "your-exotel-token";
		}

		public string execute (string from, string to, string Body)
		{
			Dictionary<string, string> postValues = new Dictionary<string, string> ();
			postValues.Add ("From", "your-agent-number");
			postValues.Add ("To", "your-customer-number");
			postValues.Add ("Body", "your-message");

			String postString = "";

			foreach (KeyValuePair<string, string> postValue in postValues) {
				postString += postValue.Key+ "=" + WebUtility.UrlEncode (postValue.Value) + "&";
			}
			postString = postString.TrimEnd ('&');
		
			ServicePointManager.ServerCertificateValidationCallback = delegate {
				return true;
			};
			string smsURL = "http://api.exotel.in/v1/Accounts/<exotel_sid>/Sms/send";
			HttpWebRequest objRequest = (HttpWebRequest)WebRequest.Create (smsURL);
			objRequest.Credentials = new NetworkCredential (this.key, this.token);
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
				postResponse = responseStream.ReadToEnd ();
				responseStream.Close ();
			}

			return (postResponse);
		}

		public static void Main (string[] args)
		{
			SendSMS s = new SendSMS ("exotel_sid","exotel_key", "exotel_token");
			string response = s.execute ("your Exophone number", "customer no.", "body");
			Console.WriteLine (response);
		}
	}
}
	



