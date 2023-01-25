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
)





// this package exist because trying to bind to popular port could terminate the program if it's been used by another program, this progam bruteforce open port then save it in $HOME/.workspace/port/servicename
//p := randomserviceport.GetPersistentRandomPort("tcp://127.0.0.1:443",".workspace", "443")
//p := randomserviceport.GetPersistentRandomPort("udp://127.0.0.1:53",".workspace, "nameserver")

//usage = PROXYPORT=`cat $HOME/.workspace/port/443` some program that need the port


// get persistent port
func GetPersistentRandomPort(portstring string, workSpace string, serviceName string) (string) {

   servicePort, _ :=  LoadPortFromWorkDir(workSpace, serviceName)

    if servicePort == ""{

      servicePort = GetRandomFreePort(portstring)
      SavePortToWorkDir(servicePort, workSpace , serviceName, true)

    }


  return servicePort
}



// get random free port
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
//this packet query if dns is available on the given port
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
  workSpace = filepath.Join(workDir, "port")
		if _, err := os.Stat( workSpace); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(workSpace, os.ModePerm)
			if err != nil {
//				log.Println(err)
			}
		}

	servicepath := filepath.Join(workSpace,  serviceName)


	var tmpFile *os.File
	if overWrite{
		tmpFile, _ = os.OpenFile(servicepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	}else{
		tmpFile, _ = os.OpenFile(servicepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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



// load save generated port

func LoadPortFromWorkDir(workSpace string, serviceName string) (string, error) {

	usr, _ := user.Current()

	workDir := filepath.Join(usr.HomeDir, workSpace)
	workSpace = filepath.Join(workDir, "port")
		if _, err := os.Stat(workSpace); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(workSpace, os.ModePerm)
			if err != nil {
//				log.Println(err)
			}
		}

	servicepath := filepath.Join(workSpace,  serviceName)

  ports, _:= os.ReadFile(servicepath)

	return string(ports), nil
}
