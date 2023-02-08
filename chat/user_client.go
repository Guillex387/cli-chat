package chat

import (
	"bufio"
	"net"
	"time"
)

// Struct to represent the client of a user
type UserClient struct {
	Conection net.Conn
}

// Opens connection to a host
func OpenConnection(ip string, port string, nickname string) (UserClient, error) {
	conn, connectError := net.Dial("tcp4", ip + ":" + port)
	if (connectError != nil) {
		return UserClient{nil}, &OpenConnectionError{}
	}
	openRequest := NewOpenInstruction(nickname) 
	conn.Write(openRequest.Bytes())
	response, writeError := bufio.NewReader(conn).ReadString('\n')
	if writeError != nil {
		return UserClient{nil}, &ConnectionIOError{}
	}
	responseInstruction := BytesToInstruction([]byte(response))
	if responseInstruction.Id == "open" {
		return UserClient{Conection: conn}, nil
	}
	return UserClient{nil}, &OpenConnectionError{}
}

// Executes a callback when receive/send a message
func (c *UserClient) MessageListen(callback func(instruction Instruction)) {
	reader := bufio.NewReader(c.Conection)
	for {
		instructionStr, _ := reader.ReadString('\n')
		instruction := BytesToInstruction([]byte(instructionStr))
		callback(instruction)
		time.Sleep(500 * time.Millisecond)
	}
}

// Send a instruction to the host
func (c *UserClient) SendInstruction(instruction Instruction) error {
	_, writeError := c.Conection.Write(instruction.Bytes())
	if (writeError != nil) {
		return &ConnectionIOError{}
	}
	return nil
}

// Closes the connection to host
func (c *UserClient) Close() {
	c.SendInstruction(NewEndInstruction())
	c.Conection.Close()
	c.Conection = nil
}
