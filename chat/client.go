package chat

import "cli-chat/ins"

// A interface for interact with clients
type Client interface {
  Event() *Event
  SendInstruction(ins.Instruction) error
  Listen()
  Close()
}
