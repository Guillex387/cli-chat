package messages

var instructionsIds = map[string]int {
	"": 0,       // Send/Receive normal message
	"end": 1,    // Close connection
	"sendf": 2,  // Send/Receive file
	"kill": 3,   // Closes a user connection
	"log": 4,    // Server logs
}

var idsInstructions = map[int]string {
	0: "",       // Send/Receive normal message
	1: "end",    // Close connection
	2: "sendf",  // Send/Receive file
	3: "kill",   // Closes a user connection
	4: "log",    // Server logs
}

type Message struct {
	Instruction int
	Body string
	Input bool
}

// Parse message inputs to a instructions
func MessageParse(msg string, input bool) Message {
	instructionId := ""
	body := ""
	readingInstruction := msg[0] == '/'
	for i := 1; i < len(msg); i++ {
		if readingInstruction {
			if msg[i] == ' ' {
				readingInstruction = false
				continue
			}
			instructionId += string(msg[i])
		} else {
			body += string(msg[i])
		}
	}
	return Message {
		Instruction: instructionsIds[instructionId],
		Body: body,
		Input: input,
	}
}

func MessageStringify(msg Message) string {
	return idsInstructions[msg.Instruction] + msg.Body
}