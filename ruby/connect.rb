require './config'
require_relative './net_http'

class Connect
    def initialize(url, params)
        @url = url
        @params = params
    end

    def connect_to_flow
        NetHTTP.post(@url, @params)
    end
    
    def connect_to_agent
        NetHTTP.post(@url, @params)
    end

    def connect_to_sms
        NetHTTP.post(@url, @params)
    end
end