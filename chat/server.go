package chat

import (
	"bufio"
	"net"
)

type Server struct {
	Listener net.Listener
	UserArray []User 
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
		if user.Name != exception {
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
		// TODO: Adds the instruction switch
		}
	}
}

func (s *Server) Adduser(user User) {
	s.UserArray = append(s.UserArray, user)
	userIndex := len(s.UserArray) - 1
	go s.ListenUser(userIndex)
	s.ReplyInstruction(Instruction {Id: "log", Body: user.Name + " added to server"}, user.Name)
}

func (s *Server) Listen() {
	for {
		conn, _ := s.Listener.Accept()
		instruction_str, _ := bufio.NewReader(conn).ReadString('\n')
		instruction := InstructionParse(instruction_str)
		if instruction.Id == "open" {
			userName := instruction.Body
			if s.FindUser(userName) != -1 {
				instruction := Instruction {Id: "error", Body: "The name already exists"}
				conn.Write([]byte(instruction.String()))
				conn.Close()
			} else {
				s.Adduser(User {Name: userName, Conection: conn})
			}
		} else {
			instruction := Instruction {Id: "error", Body: "Your father works in Colombia"}
			conn.Write([]byte(instruction.String()))
			conn.Close()
		}
	}
}