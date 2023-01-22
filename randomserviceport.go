package randomserviceport

import
(
  "github.com/antelman107/net-wait-go/wait"
  "time"
  "math/rand"
  "strings"
  "net/url"
  "net"
  "fmt"
)

func GetRandomFreePort(portstring string) (string) {

  OpenPort :=	OpenPortAny(portstring)

 u, _ := url.Parse(portstring)
host, port, _ := net.SplitHostPort(u.Host)

  if(OpenPort){
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)


  for {
    port = string(fmt.Sprintf("%v", r1.Intn(19999)))
    portstring = u.Scheme + "://" + host + ":" + port
    OpenPort :=	OpenPortAny(portstring)
    if !OpenPort{
      break
    }
  }



  }

  return  port

}



func OpenPortAny( portstring string) (bool) {
  connected := true

 u, _ := url.Parse(portstring)
// host, port, _ := net.SplitHostPort(u.Host)



//hostport := host + ":" + port
switch strings.ToLower(u.Scheme) {
case "udp":

  e := wait.New(
    wait.WithProto("udp"),
    wait.WithUDPPacket([]byte{0x00,  0x01,  0x01,  0x00,  0x00,  0x01,  0x00,  0x00,  0x00,  0x00,  0x00,  0x00,  0x06,  0x67,  0x6F,  0x6F,  0x67,  0x6C,  0x65,  0x03,  0x63,  0x6F,  0x6D,  0x00,  0x00,  0x0F,  0x00,  0x01, 0x01}),
//  wait.WithDebug(true),
    wait.WithDeadline(time.Second*2),
  )
  if !e.Do([]string{ u.Host }) {
    connected = false
//  return
  }
case "tcp":
if !wait.New(
  wait.WithProto("tcp"),
  wait.WithDeadline(2*time.Second),
//wait.WithDebug(true),
).Do([]string{ u.Host }) {
  connected = false
}
default:
  if !wait.New(
    wait.WithProto("tcp"),
    wait.WithDeadline(2*time.Second),
//  wait.WithDebug(true),
  ).Do([]string{u.Host}) {
    connected = false
  }

}


  return   connected

}
