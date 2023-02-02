package chat

import (
	"encoding/base64"
	"strings"
)

// Instructions:
//
// "" -> Send a normal message
//
// "open" -> Open a client connection
//
// "end" -> Close a client/server connection
//
// "sendf" -> Send a file
//
// "kill" -> kill a user (Host only)
//
// "log" -> Send a server log (Host only)
type Instruction struct {
  Id string
  Args [][]byte
}

// The constructor of Instruction
func NewIntruction(id string, args ...[]byte) Instruction {
  return Instruction {Id: id, Args: args};
}

// Parse instructions to string
func (i Instruction) Bytes() []byte {
  result := i.Id
  for _, data := range i.Args {
    result += " " + base64.StdEncoding.EncodeToString(data)
  }
  return []byte(result)
}

// Parse instrcution inputs to a instructions
func InstructionParse(data string) Instruction {
  dataSec := strings.Split(data[:len(data) - 1], " ")
  args := make([][]byte, len(dataSec) - 1)
  for i, _ := range args {
    args[i], _ = base64.StdEncoding.DecodeString(dataSec[i + 1])
  }
  return NewIntruction(dataSec[0], args...)
}

// Creates a log instruction
func NewlogInstruction(log string) Instruction {
  return NewIntruction("log", []byte(log))
}

// Creates a kill instruction
func NewKillInstruction(userName string) Instruction {
  return NewIntruction("kill", []byte(userName))
}

// Creates a error instruction
func NewErrorInstruction(error string) Instruction {
  return NewIntruction("error", []byte(error))
}

// Creates a end instruction
func NewEndInstruction() Instruction {
  return NewIntruction("end")
}

// Creates a open instruction
func NewOpenInstruction(userName string) Instruction {
  return NewIntruction("open", []byte(userName))
}

// Creates a message instruction
func NewMsgInstruction(userName string, msg string) Instruction {
  return NewIntruction("", []byte(userName), []byte(msg))
}
