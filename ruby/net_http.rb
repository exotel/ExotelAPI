require 'uri'
require 'net/http'

class NetHTTP

  def self.post(url, params)
    res = Net::HTTP.post_form(URI(url), params)
    puts res.read_body
  rescue TimeoutError => te
    puts te
  rescue StandardError => se
    puts se
  end
end