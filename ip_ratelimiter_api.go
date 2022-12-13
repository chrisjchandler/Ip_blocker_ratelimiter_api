package main

import (
"fmt"
"log"
"net/http"
"os/exec"
)

func rateLimit(w http.ResponseWriter, r *http.Request) {
// Rate limit DNS queries to 100 requests per second
//ip := r.URL.Query().Get("ip")
cmd := exec.Command("iptables", "-A", "INPUT", "-p", "udp", "--dport", "53", "-m", "limit", "--limit", "100/sec", "-j", "ACCEPT")
err := cmd.Run()
if err != nil {
log.Fatal(err)
}
fmt.Fprintln(w, "Rate limiting DNS queries to 100 requests per second")
}

func filterQueries(w http.ResponseWriter, r *http.Request) {
// Block DNS queries from specified IP addresses
ip := r.URL.Query().Get("ip")
cmd := exec.Command("iptables", "-A", "INPUT", "-s", ip, "-p", "udp", "--dport", "53", "-j", "DROP")
err := cmd.Run()
if err != nil {
log.Fatal(err)
}
fmt.Fprintln(w, "Blocking DNS queries from specified IP addresses")
}

func main() {
http.HandleFunc("/ratelimit", rateLimit)
http.HandleFunc("/filterqueries", filterQueries)
log.Fatal(http.ListenAndServe(":8080", nil))
}

//To use this API, you can make curl requests to the /ratelimit and /filterqueries endpoints, such as:

//$ curl http://localhost:8080/ratelimit
//Rate limiting DNS queries to 100 requests per second

//$ curl http://localhost:8080/filterqueries
//Blocking DNS queries from specified IP addresses
