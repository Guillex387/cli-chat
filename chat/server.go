package chat

import (
	"bufio"
	"cli-chat/ins"
	"net"
	"time"
)

// A struct for manage the chat hosting
type Server struct {
  ConnListener net.Listener
  UserArray []User
  SendEvent Event
  Listener Listener
}

// Validates if the nickname of a user
// isn't a reserved word
func ValidName(name string) bool {
  switch name {
    case "Server":
      return false
    case "log":
      return false
    case "error":
      return false
    case "You":
      return false
    default:
      return true
  }
}

// Creates a server
func InitServer(port string) (Server, error) {
  connListener, err := net.Listen("tcp", ":" + port)
  if err != nil {
    return Server{nil, nil, NewEvent(), NewListener()}, OpenServerError{}
  }
  return Server{
    ConnListener: connListener,
    UserArray: make([]User, 0),
    SendEvent: NewEvent(),
    Listener: NewListener(),
  }, nil
}

// Send a instruction to all user (with an exception)
func (s *Server) ReplyInstruction(instruction ins.Instruction, exception string) {
  for _, user := range s.UserArray {
    if exception == "" || user.Name != exception {
      user.SendInstruction(instruction)
    }
  }
  s.SendEvent.Trigger(instruction)
}

// Find the index of a user in the UserArray
// return -1 if not found
func (s *Server) FindUser(name string) int {
  for i, user := range s.UserArray {
    if user.Name == name {
      return i
    }
  }
  return -1
}

// Adds a user to the chat
func (s *Server) AddUser(user User) {
  s.UserArray = append(s.UserArray, user)
  s.ReplyInstruction(ins.NewLogInstruction(user.Name + " joined to chat"), "")
  user.Listen()
  user.MessageEvent.OnAny(func(this EventListener, instruction ins.Instruction) {
    s.ManageUserInstruction(&user, instruction)
  })
}

// Removes a user from the chat
func (s *Server) DeleteUser(user User) {
  findIndex := s.FindUser(user.Name)
  if findIndex == -1 {
    return
  }
  s.UserArray = append(s.UserArray[0:findIndex], s.UserArray[(findIndex + 1):]...)
  user.Close()
  s.ReplyInstruction(ins.NewLogInstruction(user.Name + " closed connection"), "")
}

// Remove all users from the chat
func (s *Server) DeleteAllUsers() {
  for _, user := range s.UserArray {
    user.Close()
    s.ReplyInstruction(ins.NewLogInstruction(user.Name + " closed connection"), "")
  }
}

// Listen to new user connections
func (s *Server) Listen() {
  s.Listener.Open(func(stop chan struct{}) {
    for {
      // Stop signal
      select {
        case <- stop:
          return
        default:
          break
      }

      conn, err := s.ConnListener.Accept()
      if err != nil {
        continue
      }
      instruction_str, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        continue
      }
      instruction := ins.BytesToInstruction([]byte(instruction_str)) 

      if instruction.Id != "open" {
        conn.Write(ins.NewErrorInstruction("Unknow instruction").Bytes())
        conn.Close()
        continue
      }

      userName := string(instruction.Args[0])
      if !ValidName(userName) {
        conn.Write(ins.NewErrorInstruction("The name is not valid").Bytes())
        conn.Close()
        continue
      }

      if s.FindUser(userName) != -1 {
        conn.Write(ins.NewErrorInstruction("The name already exists").Bytes())
        conn.Close()
        continue
      }

      user := NewUser(userName, conn)
      user.SendInstruction(ins.NewOkInstruction())
      s.AddUser(user)
      time.Sleep(500 * time.Millisecond)
    }
  })
}

// Close the server
func (s *Server) Close() {
  s.DeleteAllUsers()
  s.SendEvent.Clear()
  s.ConnListener.Close()
  s.Listener.Close()
}
