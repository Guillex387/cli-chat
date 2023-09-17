package chat

import (
	"cli-chat/ins"
	"os/exec"
	"strings"
)

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
  commandName := string(instruction.Args[0])
  strArg := make([]string, 0)
  for _, arg := range instruction.Args[1:] {
    strArg = append(strArg, string(arg))
  }

  cmd := exec.Command(commandName, strArg...)
  output, err := cmd.Output()

  if err != nil {
    c.ReceiveEvent.Trigger(ins.NewErrorInstruction(err.Error()))
    return
  }
  msg := commandName + " " + strings.Join(strArg, " ") + "\n" + string(output)
  c.SendInstruction(ins.NewMsgInstruction("", msg))
}
