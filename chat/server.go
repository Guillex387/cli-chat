package chat

import (
	"bufio"
	"net"
	"time"
)

// A struct for manage the chat hosting
type Server struct {
	Listener net.Listener
	UserArray []User
  SendEvent Event
}

// Validates if the nickname of a user
// isn't a reserved word
func ValidName(name string) bool {
	reservedNames := map[string]int {
		"Server": 1,
		"log": 1,
		"error": 1,
	}
	return (reservedNames[name] != 1)
}

// Creates a server
func InitServer(port string) (Server, error) {
	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return Server{nil, nil, NewEvent()}, &OpenServerError{}
	}
	return Server {Listener: listener, UserArray: make([]User, 0)}, nil
}

// Send a instruction to all user (with an exception)
func (s *Server) ReplyInstruction(instruction Instruction, exception string) {
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
	go user.Listen()
  user.MessageListen(func(instruction Instruction) {
    switch instruction.Id {
		  case "":
			  s.ReplyInstruction(instruction, user.Name)
		  case "end":
			  s.DeleteUser(user)
		  // case "sendf":
			  // TODO: define this feature
		  default:
			  user.SendInstruction(NewErrorInstruction("Unknow instruction"))
		}
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
	s.ReplyInstruction(NewlogInstruction(user.Name + " closed connection"), "")
}

// Listen to new user connections
func (s *Server) Listen() {
	for {
    time.Sleep(500 * time.Millisecond)
		conn, _ := s.Listener.Accept()
		instruction_str, _ := bufio.NewReader(conn).ReadString('\n')
		instruction := BytesToInstruction([]byte(instruction_str)) 
		if instruction.Id != "open" {
      conn.Write(NewErrorInstruction("Unknow instruction").Bytes())
			conn.Close()
      continue
    }
		userName := string(instruction.Args[0])
		if !ValidName(userName) {
			conn.Write(NewErrorInstruction("The name is not valid").Bytes())
			conn.Close()
      continue
		}
    if s.FindUser(userName) != -1 {
			conn.Write(NewErrorInstruction("The name already exists").Bytes())
			conn.Close()
      continue
		}
    user := NewUser(userName, conn)
    s.AddUser(user)
		user.SendInstruction(NewOkInstruction())
	}
}
