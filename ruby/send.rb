require './config'
require_relative './connect.rb'
include  Cred

Connect.new(Cred.sms,     'To' => 'your-agent-number', 'From' => 'your-customer-number', 'Body' => 'your-message').connect_to_sms
Connect.new(Cred.connect, 'To' => 'your-agent-number',  'From' => 'your-customer-number', 'Url' => 'http://my.exotel.com/<exotel-sid>/exoml/start_voice/<app_id>').connect_to_flow
Connect.new(Cred.connect, 'To' => 'your-agent-number', 'From' => 'your-customer-number').connect_to_agent
