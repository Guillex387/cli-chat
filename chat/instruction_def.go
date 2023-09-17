package chat

import (
	"cli-chat/ins"
	"fmt"
	"os/exec"
	"strings"
)

// Manage the server instructions behavior
func (s *Server) ManageServerInstruction(instruction ins.Instruction) {
  switch instruction.Id {
    case "":
      s.ManageMsg(instruction)
    case "kill":
      s.ManageKill(instruction)
    case "end":
      s.ManageEnd()
    case "users":
      s.ManageUsers()
    case "cmd":
      s.ManageCmd(instruction)
    case "clear":
      s.SendEvent.Trigger(instruction)
    default:
      error := ins.NewErrorInstruction("Unknown instruction")
      s.SendEvent.Trigger(error)
  }
}

// Manage the user instruction behavior (in the server)
func (s *Server) ManageUserInstruction(user *User, instruction ins.Instruction) {
  switch instruction.Id {
    case "":
      s.ManageUserMsg(user, instruction)
    case "end":
      s.ManageUserEnd(user)
    case "users":
      s.ManageUserUsers(user)
    default:
      error := ins.NewErrorInstruction("Unknown instruction")
      user.SendInstruction(error)
  }
}

// Server instructions

// Manage server message instruction
func (s *Server) ManageMsg(instruction ins.Instruction) {
  instruction.Args[0] = []byte("Server")
  s.ReplyInstruction(instruction, "")
}

// Manage server kill instruction
func (s *Server) ManageKill(instruction ins.Instruction) {
  user := s.FindUser(string(instruction.Args[0]))
  if user == -1 {
    errorMsg := fmt.Sprintf("User '%s' not found", string(instruction.Args[0]))
    s.SendEvent.Trigger(ins.NewErrorInstruction(errorMsg))
    return
  }
  s.DeleteUser(s.UserArray[user])
}

// Manage server end instruction
func (s *Server) ManageEnd() {
  s.SendEvent.Trigger(ins.NewEndInstruction())
  s.Close()
}

// Manage server users instruction
func (s *Server) ManageUsers() {
  users := "User list:"
  for _, user := range s.UserArray {
    users += " " + user.Name
  }
  s.SendEvent.Trigger(ins.NewLogInstruction(users))
}

// Manage server cmd instruction
func (s *Server) ManageCmd(instruction ins.Instruction) {
  command_name := string(instruction.Args[0])
  strArg := make([]string, 0)
  for _, arg := range instruction.Args[1:] {
    strArg = append(strArg, string(arg))
  }

  cmd := exec.Command(command_name, strArg...)
  output, err := cmd.Output()

  if err != nil {
    s.SendEvent.Trigger(ins.NewErrorInstruction(err.Error()))
    return
  }
  msg := command_name + " " + strings.Join(strArg, " ") + "\n" + string(output)
  s.ManageMsg(ins.NewMsgInstruction("", msg))
}

// User instructions

// Manage user message instruction
func (s *Server) ManageUserMsg(user *User, instruction ins.Instruction) {
  instruction.Args[0] = []byte(user.Name)
  s.ReplyInstruction(instruction, user.Name)
}

// Manage user end instruction
func (s *Server) ManageUserEnd(user *User) {
  s.DeleteUser(*user)
}

// Manage user users instruction
func (s *Server) ManageUserUsers(user *User) {
  users := "User list:"
  for _, user := range s.UserArray {
    users += " " + user.Name
  }
  user.SendInstruction(ins.NewLogInstruction(users))
}
