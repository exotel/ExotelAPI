# EXOTEL-C#

# DOCUMENTATION:-

The documentation for the Exotel API can be found https://developer.exotel.com/api/#intro.

# SUPPORTED C# VERSIONS

    This library supports the following C# implementations: C# 1.0 .NET Framework 1.0/1.1
            C# 2.0 .NET Framework 2.0
            C# 3.0 .NET Framework 3.0\3.5
            C# 4.0 .NET Framework 4.0
            C# 5.0 .NET Framework 4.5
            C# 6.0 .NET Framework 4.6
            C# 7.0 .NET Core 2.0
            C# 8.0 .NET Core 3.0

# INSTALLATION

    At Ubuntu terminal, download and install following packages via terminal:
    $ sudo apt-add-repository ppa:directhex/ppa
    $ sudo apt-get update
    $ sudo apt-get install monodevelop
    Incase you face any issue try running below command
    $ sudo apt-get -f install
    Now, install mcs utility which can compile your C# into exe file
    $ sudo apt-get install mcs

# PUT  YOUR OWN CREDENTIAL HERE

    account_sid = 'ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' auth_token = 'yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy'
    

# MAKE A CALL

```csharp
using System; 
using System.Collections.Generic;
using System.IO; 
using System.Net;

namespace ExotelSDK { public class Connects { private string SID = null; private string KEY = null; private string TOKEN = null ;

    public Connects (string SID, string KEY, string TOKEN)
    {
        this.SID   = "your-exotel-sid";
        this.KEY   = "your-exotel-key";
        this.TOKEN = "your-exotel-token";
    }

    public string connectCustomerToAgent (string from, string to, string callerID,
                                          string callType, string timeLimit = null,
                                          string timeOut = null, string statusCallback = null)
    {
        Dictionary<string, string> postValues = new Dictionary<string, string> ();
        postValues.Add ("From", "your-agent-number");
        postValues.Add ("To", "your-customer-number");
        postValues.Add ("CallerID", "your-callerid");
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
        string smsURL = "https://api.exotel.in/v1/Accounts/<exotel_sid>/Calls/connect";
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
        Connects c = new Connects ("exotel_sid","exotel_key","exotel-token");
        string response = c.connectCustomerToAgent("From", "To", "your exotel VN", "trans");
    Console.WriteLine(response);
    }
}
}
```

# SEND AN SMS

```csharp
package main
using System; using System.Collections.Generic; using System.IO;
using System.Net;
namespace ExotelSDK { public class SendSMS { private string SID = null; private string key = null; private string token =null ;

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

```
# EXECUTING C# CODE

    For building and executing c# code (tested on linux Mono):- 
    1) mcs filename.cs .
    2) mono filename.exe.

# RESULT

    {"Call":{"Sid":"40ccc7656e1a3eb66c6cb2be57d1142k",
    "ParentCallSid":null, 
    "DateCreated":"2020-02-20 18:42:19", 
    "DateUpdated":"2020-02-20 18:42:19", 
    "AccountSid":"Exotel", 
    "To":"+919890000000", 
    "From":"+917800000000", 
    "PhoneNumberSid":"+60327712799", 
    "Status":"in-progress", 
    "StartTime":"2020-02-20 18:42:19", 
    "EndTime":null, 
    "Duration":null, 
    "Price":null,
    "Direction":"outbound-api",
    "AnsweredBy":null,
    "ForwardedFrom":null, 
    "CallerName":null, "Uri":"\/v1\/Accounts\/Exotel\/Calls\/40ccc7656e1a3eb66c6cb2be57d1142k.json", 
    "RecordingUrl":null, 
    "SubResourceUris":{"Legs":"\/v1\/Accounts\/Exotel\/Calls\/40ccc7656e1a3eb66c6cb2be57d1142k.json\/Legs"}}}
    
  
