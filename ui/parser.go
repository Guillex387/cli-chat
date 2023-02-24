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
  input += " "
  splits := make([]string, 0)
  openSemiCol := false
  token := ""
  // This loops checks the letters and
  // separates the input in tokens
  for _, letter := range input[1:] {
    if letter == '"' {
      if openSemiCol {
        splits = append(splits, token)
        openSemiCol = false
        token = ""
        continue
      }
      openSemiCol = true
      token = ""
      continue
    }
    if letter == ' ' && !openSemiCol {
      if len(token) != 0 {
        splits = append(splits, token)
      }
      token = ""
      continue
    }
    token += string(letter)
  }
  // Checks if exists a remain token
  if len(token) != 0 {
    splits = append(splits, token)
  }
  return InputInstruction{
    Id: splits[0],
    Body: splits[1:],
  }
}
