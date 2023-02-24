package chat

import (
  "bufio"
  "net"
  "time"
)

// Struct to represent the client of a user
type UserClient struct {
  Conection net.Conn
  ReceiveEvent Event
}

// Opens connection to a host
func OpenConnection(ip string, port string, nickname string) (Client, error) {
  conn, connectError := net.Dial("tcp4", ip + ":" + port)
  if (connectError != nil) {
    return nil, OpenConnectionError{}
  }
  openRequest := NewOpenInstruction(nickname) 
  conn.Write(openRequest.Bytes())
  response, writeError := bufio.NewReader(conn).ReadString('\n')
  if writeError != nil {
    return nil, ConnectionIOError{}
  }
  responseInstruction := BytesToInstruction([]byte(response))
  if responseInstruction.Id == "ok" {
    return &UserClient{Conection: conn, ReceiveEvent: NewEvent()}, nil
  }
  return nil, OpenConnectionError{}
}

// Executes a callback when receive/send a message
func (c *UserClient) Listen() {
  reader := bufio.NewReader(c.Conection)
  for {
    instructionStr, _ := reader.ReadString('\n')
    instruction := BytesToInstruction([]byte(instructionStr))
    c.ReceiveEvent.Trigger(instruction)
    time.Sleep(500 * time.Millisecond)
  }
}

// Creates a message listener
func (c *UserClient) MessageListen(callback func(instruction Instruction)) func() {
  return c.ReceiveEvent.CreateListener(callback)
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
  c.ReceiveEvent.Clear()
}
