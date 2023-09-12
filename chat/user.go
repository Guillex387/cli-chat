package chat

import (
	"bufio"
	"net"
	"time"
  "cli-chat/ins"
)

// User (in server) struct representation
type User struct {
  Name string
  Conection net.Conn
  MessageEvent Event
  Listener Listener
}

// Creates a new user
func NewUser(name string, connection net.Conn) User {
  return User{Name: name, Conection: connection, MessageEvent: NewEvent(), Listener: NewListener()}
}

// Send a instruction to the user
func (u *User) SendInstruction(instruction ins.Instruction) {
  u.Conection.Write(instruction.Bytes())
}

// Listen the incoming message from the user connection
func (u *User) Listen() {
  u.Listener.Open(func(stop chan struct{}) {
    reader := bufio.NewReader(u.Conection)
    for {
      // Stop signal
      select {
        case <- stop:
          return
        default:
          break
      }

      time.Sleep(500 * time.Millisecond)
      instruction_str, err := reader.ReadString('\n')
      if err != nil {
        continue
      }
      instruction := ins.BytesToInstruction([]byte(instruction_str))
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
  u.Listener.Close()
  u.SendInstruction(ins.NewlogInstruction("The host kill you"))
  u.SendInstruction(ins.NewEndInstruction())
  u.MessageEvent.Clear()
  u.Conection.Close()
}
