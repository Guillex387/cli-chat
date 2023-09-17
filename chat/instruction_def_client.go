package chat

import "cli-chat/ins"

// Manage the client instructions behavior
func (c *UserClient) ManageInstruction(instruction ins.Instruction) {
  switch instruction.Id {
    case "cmd":
      c.ManageCmd(instruction)
    case "clear":
      c.ReceiveEvent.Trigger(instruction)
    default:
      if instruction.Id == "" {
        instruction.Args[0] = []byte("You")
      }
      _, writeErr := c.Connection.Write(instruction.Bytes())
      if writeErr != nil {
        errMsg := ConnectionIOError{}.Error()
        c.ReceiveEvent.Trigger(ins.NewErrorInstruction(errMsg))
        break
      }
      c.ReceiveEvent.Trigger(instruction)
  }
}

// Manage cmd instruction
func (c *UserClient) ManageCmd(instruction ins.Instruction) {
  outInstruction := ExecuteCmd(instruction)
  if outInstruction.Id == "error" {
    c.ReceiveEvent.Trigger(outInstruction)
    return
  }
  c.SendInstruction(outInstruction)
}
