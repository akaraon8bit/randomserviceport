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
	"os/user"
	"path/filepath"
	"os"
	"errors"
	"strconv"
)


func GetRandomFreePort(portstring string) (int) {

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
  x, _ := strconv.Atoi(port)

  return  x

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



// save generated port

func SavePortToWorkDir(texts string, workSpace string, serviceName string, overWrite bool) (string, error) {

	usr, _ := user.Current()

	workDir := filepath.Join(usr.HomeDir, workSpace)
	resultpath := filepath.Join(workDir, "port")
		if _, err := os.Stat(resultpath); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(resultpath, os.ModePerm)
			if err != nil {
//				log.Println(err)
			}
		}

	resultfilepath := filepath.Join(resultpath,  serviceName)


	var tmpFile *os.File
	if overWrite{
		tmpFile, _ = os.OpenFile(resultfilepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	}else{
		tmpFile, _ = os.OpenFile(resultfilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}


	text := []byte(texts)
	if _, err := tmpFile.Write(text); err != nil {
		return "error", err
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		return "error", err
	}

	return tmpFile.Name(), nil
}
