package chat

import (
	"net"
)

// User struct representation
type User struct {
  Name string
  Conection net.Conn
}

// Send a instruction to the user
func (c *User) SendInstruction(instruction Instruction) {
  c.Conection.Write(instruction.Bytes())
}
