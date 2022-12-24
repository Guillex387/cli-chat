package chat

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
  Body string
}


func NewIntruction(id string, body string) Instruction {
  return Instruction {Id: id, Body: body};
}

// Parse instructions to string messages
func (i Instruction) String() string {
  return i.Id + i.Body + "\n"
}

// Parse instrcution inputs to a instructions
func InstructionParse(data string) Instruction {
  input := RemoveLinebreaks(data)
  inputInstruction := ""
  body := ""
  // Check if message contents a instruction
  readingInstruction := input[0] == '/'
  // Message parser
  for i := 1; i < len(input); i++ {
    if readingInstruction {
      if input[i] == ' ' {
        readingInstruction = false
        continue
      }
      inputInstruction += string(input[i])
    } else {
      body += string(input[i])
    }
  }
  return Instruction {
    Id: inputInstruction,
    Body: body,
  }
}