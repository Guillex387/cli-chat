package chat

import (
	"net"
)

type User struct {
  Name string
  Conection net.Conn
}

func (c *User) SendInstruction(instruction Instruction) {
  c.Conection.Write([]byte(instruction.String()))
}