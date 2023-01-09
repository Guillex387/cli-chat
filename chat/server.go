package chat

import (
	"bufio"
	"net"
)

type Server struct {
	Listener net.Listener
	UserArray []User 
}

func ValidName(name string) bool {
	reservedNames := map[string]int {
		"server": 1,
		"log": 1,
		"error": 1,
	}
	return (reservedNames[name] != 1)
}

func InitServer(port string) Server {
	listener, _ := net.Listen("tcp", ":" + port)
	return Server {Listener: listener, UserArray: make([]User, 0)}
}

func (s *Server) FindUser(name string) int {
	for i, user := range s.UserArray {
		if user.Name == name {
			return i
		}
	}
	return -1
}

func (s *Server) ReplyInstruction(instruction Instruction, exception string) {
	for _, user := range s.UserArray {
		if exception == "" || user.Name != exception {
			user.SendInstruction(instruction)
		}
	}
}

func (s *Server) ListenUser(id int) {
	user := s.UserArray[id]
	reader := bufio.NewReader(user.Conection)
	for {
		instruction_str, _ := reader.ReadString('\n')
		instruction := InstructionParse(instruction_str)
		switch instruction.Id {
		case "":
			s.ReplyInstruction(instruction, user.Name)
		case "end":
			s.DeleteUser(user)
			return
		// case "sendf":
			// TODO: define this feature
		default:
			user.SendInstruction(NewIntruction("error", "Unknow instruction"))
		}
	}
}

func (s *Server) AddUser(user User) {
	s.UserArray = append(s.UserArray, user)
	userIndex := len(s.UserArray) - 1
	go s.ListenUser(userIndex)
}

func (s *Server) DeleteUser(user User) {
	findIndex := s.FindUser(user.Name)
  if findIndex == -1 {
    return
  }
  slice1 := s.UserArray[0:findIndex]
  slice2 := s.UserArray[(findIndex + 1):]
  s.UserArray = append(slice1, slice2...)
  user.Conection.Close()
	s.ReplyInstruction(NewIntruction("log", user.Name + " closed connection"), "")
}

func (s *Server) Listen() {
	for {
		conn, _ := s.Listener.Accept()
		instruction_str, _ := bufio.NewReader(conn).ReadString('\n')
		instruction := InstructionParse(instruction_str)
		if instruction.Id == "open" {
			userName := instruction.Body
			if ValidName(userName) {
				conn.Write([]byte(NewIntruction("error", "The name is not valid").String()))
				conn.Close()
			} else if s.FindUser(userName) != -1 {
				conn.Write([]byte(NewIntruction("error", "The name already exists").String()))
				conn.Close()
			} else {
				s.AddUser(User {Name: userName, Conection: conn})
			}
		} else {
			conn.Write([]byte(NewIntruction("error", "Unknow instruction").String()))
			conn.Close()
		}
	}
}