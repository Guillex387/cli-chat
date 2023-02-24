package ui

// Parse a instruction from the
// user input
type InputInstruction struct {
  Id string
  Body []string
}

// Parse instruction from user input 
func ParseInstruction(input string) InputInstruction {
  // Check if is a null input and returns a void instruction
  if len(input) == 0 {
    return InputInstruction{
      Id: "",
      Body: []string{""},
    }
  }
  // Checks if input not contain a instruction id
  // and returns the message id ("") and the entire
  // input as the body
  if input[0] != '/' {
    return InputInstruction{
      Id: "",
      Body: []string{input},
    }
  }
  tokenList := make([]string, 0)
  openQuote := false
  token := ""
  // This loops checks the letters and
  // separates the input in tokens
  for _, letter := range input[1:] {
    if letter == '"' {
      if openQuote {
        tokenList = append(tokenList, token)
        openQuote = false
        token = ""
        continue
      }
      openQuote = true
      token = ""
      continue
    }
    if letter == ' ' && !openQuote {
      if len(token) != 0 {
        tokenList = append(tokenList, token)
      }
      token = ""
      continue
    }
    token += string(letter)
  }
  // Checks if exists a remain token
  if len(token) != 0 {
    tokenList = append(tokenList, token)
  }
  return InputInstruction{
    Id: tokenList[0],
    Body: tokenList[1:],
  }
}

// TODO: make a function for convert an "InputInstruction"
//       in a "chat.Instruction"
