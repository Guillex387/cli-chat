package chat

import (
	"cli-chat/ins"
	"net"
	"os/exec"
	"strings"
)

// Returns the local ip string
func GetLocalIp() string {
  conn, err := net.Dial("udp", "8.8.8.8:80")
  if err != nil {
    return ""
  }
  defer conn.Close()
  localAddr := conn.LocalAddr().(*net.UDPAddr)
  return localAddr.IP.String()
}

// Executes a command in the machine and returns
// and output instruction, and error or message
// with the command output
func ExecuteCmd(cmd ins.Instruction) ins.Instruction {
  commandName := string(cmd.Args[0])
  strArg := make([]string, 0)
  for _, arg := range cmd.Args[1:] {
    strArg = append(strArg, string(arg))
  }

  process := exec.Command(commandName, strArg...)
  output, err := process.CombinedOutput()

  if err != nil {
    errorMsg := string(output) + err.Error()
    return ins.NewErrorInstruction(errorMsg)
  }
  
  cmdStr := commandName + " " + strings.Join(strArg, " ")
  msg := cmdStr + "\n" + string(output)
  return ins.NewMsgInstruction("", msg)
}
