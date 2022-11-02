package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

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

func Client() {
	// var ip string
	// var port string
	// fmt.Printf("Put the host IP: ")
	// fmt.Scanf("%s", &ip)
	// fmt.Printf("Put the host PORT: ")
	// fmt.Scanf("%s", &port)
	con, err := net.Dial("tcp4", "127.0.0.1:12345")
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

func Server(port string) {
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
	mode := os.Args[1]
	if mode == "server" {
		Server("12345")
	} else {
		Client()
	}
}