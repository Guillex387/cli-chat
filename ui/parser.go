package ui

import "strings"

// Parse a instruction from the
// user input
type InputInstruction struct {
  Id string
  Body string
}

// Parse instruction from user input 
func ParseInstruction(input string) InputInstruction {
  if input[0] != '/' {
    return InputInstruction{
      Id: "",
      Body: input,
    }
  }
  splits := strings.Split(input[1:], " ")
  return InputInstruction{
    Id: splits[0],
    Body: strings.Join(splits[1:], " "),
  }
}
