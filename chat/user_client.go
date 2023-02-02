package chat

import (
	"bufio"
	"net"
	"time"
)

type UserClient struct {
	Conection net.Conn
}

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
	responseInstruction := InstructionParse(response)
	if responseInstruction.Id == "open" {
		return UserClient{Conection: conn}, nil
	}
	return UserClient{nil}, &OpenConnectionError{}
}

func (c *UserClient) MessageListen(callback func(instruction Instruction)) {
	reader := bufio.NewReader(c.Conection)
	for {
		instructionStr, _ := reader.ReadString('\n')
		instruction := InstructionParse(instructionStr)
		callback(instruction)
		time.Sleep(500 * time.Millisecond)
	}
}

func (c *UserClient) SendInstruction(instruction Instruction) error {
	_, writeError := c.Conection.Write(instruction.Bytes())
	if (writeError != nil) {
		return &ConnectionIOError{}
	}
	return nil
}

func (c *UserClient) CloseConnection() {
	c.SendInstruction(NewEndInstruction())
	c.Conection.Close()
	c.Conection = nil
}
