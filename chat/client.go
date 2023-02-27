package chat

// A interface for interact with clients
type Client interface {
  Event() *Event
  SendInstruction(Instruction) error
  Close()
}
