package chat

type ServerClient struct {
  Server Server
}

func CreateServerClient(server Server) ServerClient {
  return ServerClient {Server: server}
}

func (client *ServerClient) MessageListen(callback func(instruction Instruction)) {
  // TODO: define this
}

func (client *ServerClient) SendInstruction(instruction Instruction) {
  switch instruction.Id {
  case "":
    client.Server.ReplyInstruction(instruction, "")
  case "kill":
    user := client.Server.FindUser(string(instruction.Args[0]))
    client.Server.DeleteUser(client.Server.UserArray[user])
  case "end":
    client.Close()
  default:
    return
  }
}

func (client *ServerClient) Close() {
  for _, user := range client.Server.UserArray {
    client.Server.DeleteUser(user)
  }
  client.Server.Listener.Close()
}
