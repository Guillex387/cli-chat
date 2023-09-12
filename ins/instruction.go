package ins

import (
  "strings"
  "encoding/base64"
)

// Instructions:
//
// "" -> Send a normal message
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
  result += "\n"
  return []byte(result)
}

// Parse a instruction buffer to a instruction
func BytesToInstruction(buffer []byte) Instruction {
  str_buffer := RemoveLinebreaks(string(buffer))
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
