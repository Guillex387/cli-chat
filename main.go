package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
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

func FormatString(str string) string {
	result := ""
	for _, char := range str {
		if char != '\n' {
			result += string(char)
		}
	}
	result += "\n"
	return result
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Clients Management

/*
 A function for add a user to the chat
*/
func AddUser(name string, conn net.Conn, userArr *[]Client) {
	*userArr = append(*userArr, Client{Name: name, Conection: conn})
	userIndex := len(*userArr) - 1
	go UserListener(userIndex, userArr)
	messageToSend := FormatString(name + " join to the chat")
	ReplyAll(messageToSend, userArr, name)
}

/*
 A function for find a user in the user array
*/
func FindIndex(userName string, userArr *[]Client) int {
	for i := 0; i < len(*userArr); i++ {
		if (*userArr)[i].Name == userName {
			return i
		}
	}
	return -1
}

/*
 A function for delete a user in the server
*/
func DeleteUser(user Client, userArr *[]Client) {
	findIndex := FindIndex(user.Name, userArr)
	if findIndex == -1 {
		return
	}
	slice1 := (*userArr)[0:findIndex]
	slice2 := (*userArr)[(findIndex + 1):]
	*userArr = append(slice1, slice2...)
	user.Conection.Close()
	ReplyAll(FormatString(user.Name + " closed connection"), userArr, "")
}

/*
 A function that reads the actions of a user
*/
func UserListener(userIndex int, userArr *[]Client) {
	user := (*userArr)[userIndex]
	for {
		message, _ := bufio.NewReader(user.Conection).ReadString('\n')
		if message == "END\n" {
			DeleteUser(user, userArr)
			return
		}
		messageToSend := FormatString(user.Name + ": " + message)
		ReplyAll(messageToSend, userArr, user.Name)
		time.Sleep(time.Second / 2)
	}
}

/*
 A function that listen new client connections
*/
func AcceptClients(listener net.Listener, userArr *[]Client, exit chan int) {
	for {
		conn, err := listener.Accept()
		CheckError(err)
		name, _ := bufio.NewReader(conn).ReadString('\n')
		if FindIndex(name, userArr) != -1 {
			conn.Write([]byte("Error, the name already exists\n"))
			conn.Close()
			continue
		}
		AddUser(name, conn, userArr)
		time.Sleep(time.Second / 2)
	}
}

/*
 A function for send a messages to all user saved in the user array
*/
func ReplyAll(message string, userArr *[]Client, exception string) {
	fmt.Print(message)
	for  _, compi := range *userArr {
		if compi.Name != exception {
			compi.Conection.Write([]byte(message))
		}
	}
}

/*
 A function that constantly read the messages from a connection
*/
func MessagesListener(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}

// Modes

/*
 A function for create a client for the chat
*/
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
	firstIteration := true
	for {
		if firstIteration {
			fmt.Print("Insert your nick name: ")
			go MessagesListener(con)
			firstIteration = false
		}
		msg, _ := reader.ReadString('\n')
		if msg == "\n" {
			continue
		}
		con.Write([]byte(msg))
		if msg == "END\n" {
			break
		}
	}
	con.Close()
}

/*
 A function that creates the chat host
*/
func ServerMode(port string) {
	fmt.Printf("Server IP: %s\n", GetLocalIp())
	fmt.Printf("Server PORT: %s\n", port)
	listener, err := net.Listen("tcp", ":" + port)
	CheckError(err)
	var userArr []Client
	exitChannel := make(chan int, 1)
	AcceptClients(listener, &userArr, exitChannel)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: You don't put parameters")
		os.Exit(2)
	}
	mode := os.Args[1]
	if mode == "server" {
		if len(os.Args) != 3 {
			return
		}
		ServerMode(os.Args[2])
	} else {
		ClientMode()
	}
}