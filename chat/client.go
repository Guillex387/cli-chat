package chat

import (
	"bufio"
	"net"
	"time"
)

type Client struct {
	Conection net.Conn
}

func OpenConnection(ip string, port string, nickname string) (Client, error) {
	conn, connectError := net.Dial("tcp4", ip + ":" + port)
	if (connectError != nil) {
		return Client{nil}, &OpenConnectionError{}
	}
	openRequest := NewIntruction("open", nickname)
	conn.Write(openRequest.Bytes())
	response, writeError := bufio.NewReader(conn).ReadString('\n')
	if writeError != nil {
		return Client{nil}, &ConnectionIOError{}
	}
	responseInstruction := InstructionParse(response)
	if responseInstruction.Id == "open" {
		return Client{Conection: conn}, nil
	}
	return Client{nil}, &OpenConnectionError{}
}

func (c *Client) MessageListen(callback func(instruction Instruction)) {
	reader := bufio.NewReader(c.Conection)
	for {
		instruction_str, _ := reader.ReadString('\n')
		instruction := InstructionParse(instruction_str)
		callback(instruction)
		time.Sleep(500 * time.Millisecond)
	}
}

func (c *Client) SendInstruction(instruction Instruction) error {
	_, writeError := c.Conection.Write(instruction.Bytes())
	if (writeError != nil) {
		return &ConnectionIOError{}
	}
	return nil
}

func (c *Client) CloseConnection() {
	c.SendInstruction(NewIntruction("end", ""))
	c.Conection.Close()
	c.Conection = nil
}