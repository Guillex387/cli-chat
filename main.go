package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	Name string
	Conection net.Conn
}

// Utils

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Clients Management

func AddUser(name string, conn net.Conn, userArr *[]Client) {
	*userArr = append(*userArr, Client{Name: name, Conection: conn})
}

func FindIndex(user Client, userArr *[]Client) int {
	for i := 0; i < len(*userArr); i++ {
		if (*userArr)[i].Name == user.Name {
			return i
		} 
	}
	return -1
}

func DeleteUser(user Client, userArr *[]Client) {
	findIndex := FindIndex(user, userArr)
	if findIndex == -1 {
		return
	}
	slice1 := (*userArr)[0:findIndex]
	slice2 := (*userArr)[(findIndex + 1):]
	*userArr = append(slice1, slice2...)
}

// Modes

func ClientMode() {
	var ip string
	var port string
	fmt.Printf("Put the host IP: ")
	fmt.Scanf("%s", &ip)
	fmt.Printf("Put the host PORT: ")
	fmt.Scanf("%s", &port)
	con, err := net.Dial("tcp4", ip + ":" + port)
	reader := bufio.NewReader(os.Stdin)
	CheckError(err)
	for {
		fmt.Print(">> ")
		msg, _ := reader.ReadString('\n')
		_, err = con.Write([]byte(msg))
		CheckError(err)
		reply := make([]byte, 1024)
		_, err = con.Read(reply)
		CheckError(err)
		fmt.Println(string(reply))
		if msg == "END\n" {
			break
		}
	}
	con.Close()
}

func ServerMode(port string) {
	fmt.Printf("Server IP: %s\n", GetLocalIp())
	fmt.Printf("Server PORT: %s\n", port)
	listener, err := net.Listen("tcp", ":" + port)
	CheckError(err)
	conn, err := listener.Accept()
	CheckError(err)
	for {
    message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "" {
			continue
		}
		conn.Write([]byte("Received"))
    fmt.Print("\nMessage Received: ", string(message))
  }
}

func main() {
	// mode := os.Args[1]
	// if mode == "server" {
	// 	ServerMode("12345")
	// } else {
	// 	ClientMode()
	// }
}