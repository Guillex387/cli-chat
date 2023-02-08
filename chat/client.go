package chat

// A interface for interact with clients
type Client interface {
  MessageListen(callback func(instruction Instruction))
  SendInstruction(instruction Instruction)
  Close()
}
