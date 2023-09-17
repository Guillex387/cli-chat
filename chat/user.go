package chat

import (
	"bufio"
	"cli-chat/ins"
	"net"
	"time"
)

// User (in server) struct representation
type User struct {
  Name string
  Connection net.Conn
  MessageEvent Event
  Listener Listener
  refreshTime time.Duration
}

// Creates a new user
func NewUser(name string, connection net.Conn, refreshTime time.Duration) User {
  return User{name, connection, NewEvent(), NewListener(), refreshTime}
}

// Send a instruction to the user
func (u *User) SendInstruction(instruction ins.Instruction) {
  u.Connection.Write(instruction.Bytes())
}

// Listen the incoming message from the user connection
func (u *User) Listen(server *Server) {
  u.Listener.Open(func(stop chan struct{}) {
    reader := bufio.NewReader(u.Connection)
    for {
      // Stop signal
      select {
        case <- stop:
          return
        default:
          break
      }

      time.Sleep(u.refreshTime)
      instructionStr, err := reader.ReadString('\n')
      if err != nil {
        server.DeleteUser(*u, false)
        u.MessageEvent.Clear()
        u.Connection.Close()
        return
      }
      instruction := ins.BytesToInstruction([]byte(instructionStr))
      u.MessageEvent.Trigger(instruction)
    }
  })
}

// Getter of the event manager
func (u *User) Event() Event {
  return u.MessageEvent
}

// Send a close connection messages and close the connection
func (u *User) Close() {
  u.SendInstruction(ins.NewLogInstruction("The host kill you"))
  u.SendInstruction(ins.NewEndInstruction())
  u.MessageEvent.Clear()
  u.Connection.Close()
  u.Listener.Close()
}
