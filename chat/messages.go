package chat

var instructions = []string {
  "",       // Send/Receive normal message
  "end",    // Close connection
  "sendf",  // Send/Receive file
  "kill",   // Closes a user connection
  "log",    // Server logs
  "open",   // Open connection
  "error",  // Send/Receive some error
}

func existsInstruction(inputInstruction string) bool {
  for _, instruction := range instructions {
    if inputInstruction == instruction {
      return true
    }
  }
  return false
}

type Message struct {
  Instruction string
  Body string
  Input bool
}

// Parse instructions to string messages
func (msg Message) String() string {
  return msg.Instruction + msg.Body
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
  if existsInstruction(inputInstruction) {
    return Message {
      Instruction: inputInstruction,
      Body: body,
      Input: input,
    }
  }
  return Message {
    Instruction: "Unknow", // Unknow instructions
    Body: "",
    Input: input,
  }
}

