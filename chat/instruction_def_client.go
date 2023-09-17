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
        c.ReceiveEvent.Trigger(ins.NewErrorInstruction(writeErr.Error()))
        break
      }
      c.ReceiveEvent.Trigger(instruction)
  }
}

// Manage cmd instruction
func (c *UserClient) ManageCmd(instruction ins.Instruction) {
  command_name := string(instruction.Args[0])
  str_arg := make([]string, 0)
  for _, arg := range instruction.Args[1:] {
    str_arg = append(str_arg, string(arg))
  }

  cmd := exec.Command(command_name, str_arg...)
  output, err := cmd.Output()

  if err != nil {
    c.ReceiveEvent.Trigger(ins.NewErrorInstruction(err.Error()))
    return
  }
  msg := command_name + " " + strings.Join(str_arg, " ") + "\n" + string(output)
  c.SendInstruction(ins.NewMsgInstruction("", msg))
}
