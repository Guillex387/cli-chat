package chat

import (
  "encoding/base64"
  "strings"
)

// Instructions:
//
// "" -> Send a normal message
//
// "query" -> Send a query to the user
//
// "ok" -> Send a ok response
//
// "error" -> Send an error (Host only)
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

// Parse instructions to string
func (i Instruction) Bytes() []byte {
  result := i.Id
  for _, data := range i.Args {
    result += " " + base64.StdEncoding.EncodeToString(data)
  }
  return []byte(result)
}

// Parse a instruction buffer to a instruction
func BytesToInstruction(buffer []byte) Instruction {
  str_buffer := string(buffer)
  splits := strings.Split(str_buffer, " ")
  args := make([][]byte, 0)
  for _, split := range splits[1:] {
    arg, _ := base64.StdEncoding.DecodeString(split)
    args = append(args, arg)
  }
  return NewIntruction(splits[0], args...)
}

// The constructor of Instruction
func NewIntruction(id string, args ...[]byte) Instruction {
  return Instruction {Id: id, Args: args}
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

// Creates a ok instruction
func NewOkInstruction() Instruction {
  return NewIntruction("ok")
}

// Creates a query instruction
func NewQueryInstruction(query string) Instruction {
  return NewIntruction("query", []byte(query))
}

// Creates a sendf instruction
// func NewSendfInstruction(filepath string) []Instruction {
//   // TODO: define the instruction for read the file and parsed in instruction chunks
// }
