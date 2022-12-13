package messages

type Message struct {
  Instruction int
  Body string
  Input bool
}

var instructions = []string {
  "",       // Send/Receive normal message
  "end",    // Close connection
  "sendf",  // Send/Receive file
  "kill",   // Closes a user connection
  "log",    // Server logs
  "open",   // Open connection
}

// Parse message inputs to a instructions
func MessageParse(msg string, input bool) Message {
  inputInstruction := ""
  body := ""
  // Check if message contents a instruction
  readingInstruction := msg[0] == '/'
  // Message parser
  for i := 1; i < len(msg); i++ {
    if readingInstruction {
      if msg[i] == ' ' {
        readingInstruction = false
        continue
      }
      inputInstruction += string(msg[i])
    } else {
      body += string(msg[i])
    }
  }
  // Check if input instruction exists
  for id, instruction := range instructions {
    if inputInstruction == instruction {
      return Message {
        Instruction: id,
        Body: body,
        Input: input,
      }
    }
  }
  return Message {
    Instruction: -1, // Unknow instructions
    Body: "",
    Input: input,
  }
}

// Parse instructions to string messages
func MessageStringify(msg Message) string {
  return instructions[msg.Instruction] + msg.Body
}