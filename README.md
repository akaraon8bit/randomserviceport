# randomserviceport
// this package exist because trying to bind to popular port could terminate the program if it's been used by another program, this progam bruteforce open port then save it in $HOME/.workspace/port/servicename
//p := randomserviceport.GetPersistentRandomPort("tcp://127.0.0.1:443",".workspace", "443")
//p := randomserviceport.GetPersistentRandomPort("udp://127.0.0.1:53",".workspace, "nameserver")

//usage = PROXYPORT=`cat $HOME/.workspace/port/443` some program that need the p


package main

import
(
  "fmt"
  "github.com/akaraon8bit/randomserviceport"
)


func main() {

  p := randomserviceport.GetPersistentRandomPort("tcp://127.0.0.1:443",".app", "apacheproxy")
  fmt.Println(p)
