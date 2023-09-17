package chat

import (
	"bufio"
	"cli-chat/ins"
	"fmt"
	"net"
	"time"
)

// Struct to represent the client of a user
type UserClient struct {
  Connection net.Conn
  ReceiveEvent Event
  Listener Listener
  refreshTime time.Duration
}

// Opens connection to a host
func OpenConnection(ip string, port string, nickname string, refreshTime time.Duration) (Client, error) {
  conn, connectError := net.Dial("tcp4", ip + ":" + port)
  if connectError != nil {
    return nil, OpenConnectionError{}
  }
  // Send a open request to the chat
  openRequest := ins.NewOpenInstruction(nickname) 
  conn.Write(openRequest.Bytes())
  response, writeError := bufio.NewReader(conn).ReadString('\n')
  if writeError != nil {
    return nil, ConnectionIOError{}
  }
  // Parse the response
  responseInstruction := ins.BytesToInstruction([]byte(response))
  if responseInstruction.Id == "ok" {
    return &UserClient{conn, NewEvent(), NewListener(), refreshTime}, nil
  }
  if responseInstruction.Id == "error" {
    fmt.Println(string(responseInstruction.Args[0]))
  }
  return nil, OpenConnectionError{}
}

// Executes a callback when receive/send a message
func (c *UserClient) Listen() {
  c.Listener.Open(func(stop chan struct{}) {
    reader := bufio.NewReader(c.Connection)
    for {
      // Stop signal
      select {
        case <- stop:
          return
        default:
          break
      }

      time.Sleep(c.refreshTime)
      instructionStr, err := reader.ReadString('\n')
      if err != nil {
        continue
      }
      instruction := ins.BytesToInstruction([]byte(instructionStr))
      c.ReceiveEvent.Trigger(instruction)
    }
  })
}

// Getter of the event manager
func (c *UserClient) Event() *Event {
  return &c.ReceiveEvent
}

// Send a instruction to the host
func (c *UserClient) SendInstruction(instruction ins.Instruction) {
  c.ManageInstruction(instruction)
}

// Closes the connection to host
func (c *UserClient) Close() {
  c.Listener.Close()
  c.ReceiveEvent.Clear()
  c.Connection.Close()
}
