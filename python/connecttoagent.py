import requests

from pprint import pprint

from settings import sid, token, callurl,from_no,to


def connect_customer_to_agent(
        sid,
        token,
        agent_no,
        customer_no,
        callerid,
        timelimit=None,
        timeout=None,
        calltype = 'trans'
    ):
    return requests.post(
        callurl,
        auth = (sid, token),
        data = {
            'From': from_no,
            'To'  : to,
          
        }
    )

if __name__ == '__main__':
    r = connect_customer_to_agent(
        sid,
        token,
        agent_no    = "your-agent-number",
        customer_no = "your-customer-number",
        callerid    = "<Your-Exotel-virtual-number>",
        timelimit   = "<time-in-seconds>",  
        timeout     = "<time-in-seconds>",  
        calltype    = "trans"  
    )
    print (r.status_code)
    pprint(r.json())
