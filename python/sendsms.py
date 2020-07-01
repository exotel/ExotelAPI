import requests

from pprint import pprint

from settings import sid, token, smsurl, from_no, to


def send_message(sid, token, sms_from, sms_to, sms_body):
    return requests.post(
        smsurl,
        auth=(sid, token),
        data={
            'From': from_no,
            'To': to,
            'Body': "your-message"
        }
    )


if __name__ == '__main__':
    r = send_message(
        sid,
        token,
        sms_from='your-agent-number',
        sms_to='your-customer-number',
        sms_body='your-message'
    )
    print(r.status_code)
    pprint(r.json())
