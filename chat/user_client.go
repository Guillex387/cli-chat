package chat

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// Struct to represent the client of a user
type UserClient struct {
  Conection net.Conn
  ReceiveEvent Event
  Listener Listener
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
    return &UserClient{Conection: conn, ReceiveEvent: NewEvent(), Listener: NewListener()}, nil
  }
  if responseInstruction.Id == "error" {
    fmt.Println(string(responseInstruction.Args[0]))
  }
  return nil, OpenConnectionError{}
}

// Executes a callback when receive/send a message
func (c *UserClient) Listen() {
  c.Listener.Open(func(stop chan struct{}) {
    c.ReceiveEvent.Trigger(NewlogInstruction("Opened the listener"))
    for {
      // Stop signal
      select {
        case <- stop:
          return
        default:
          break
      }

      instructionStr, _ := bufio.NewReader(c.Conection).ReadString('\n')
      instruction := BytesToInstruction([]byte(instructionStr))
      c.ReceiveEvent.Trigger(instruction)
      time.Sleep(500 * time.Millisecond)
    }
  })
}

// Getter of the event manager
func (c *UserClient) Event() *Event {
  return &c.ReceiveEvent
}

// Send a instruction to the host
func (c *UserClient) SendInstruction(instruction Instruction) error {
  _, writeError := c.Conection.Write(instruction.Bytes())
  if (writeError != nil) {
    return &ConnectionIOError{}
  }
  c.ReceiveEvent.Trigger(instruction)
  return nil
}

// Closes the connection to host
func (c *UserClient) Close() {
  c.Listener.Close()
  c.ReceiveEvent.Clear()
  c.Conection.Close()
}
