import requests

from pprint import pprint

from settings import *


def connect_customer(
        sid, 
        token,
        customer_no, 
        agent_no,
        url,
        timelimit=None,
        timeout=None,
        calltype="trans",
    ):
    return requests.post(
        callurl,
        auth=(sid, token),
        data={
            'From': from_no,
            'To'  : to,   
        }
    )


if __name__ == '__main__':
    r = connect_customer(
        sid,
        token,
        customer_no  = "your-customer-number",
        agent_no     = "your-agent-number",
        url          = "http://my.exotel.in/exoml/start/<app_id>",
        timelimit    = "<time-in-seconds>",  
        timeout      = "<time-in-seconds>",  
        calltype     = "trans",  
        )
    print (r.status_code)
    pprint(r.json())
