To use this API, you can make curl requests to the /ratelimit and /filterqueries endpoints, such as:

curl http://localhost:8080/ratelimit
Rate limiting DNS queries to 100 requests per second

curl http://localhost:8080/filterqueries
Blocking DNS queries from specified IP addresses


To pass the IP address in the curl command, you can use the query parameter syntax, such as:

curl http://localhost:8080/filterqueries?ip=192.168.1.100
Blocking DNS queries from specified IP addresses
