package chat

// Struct that represent the server client
type ServerClient struct {
  Server Server
}

// Creates a client for a server
func CreateServerClient(server Server) Client {
  return &ServerClient {Server: server}
}

// Getter of the event manager
func (c *ServerClient) Event() Event {
  return c.Server.SendEvent
}

// Send a instruction to server
func (c *ServerClient) SendInstruction(instruction Instruction) error {
  switch instruction.Id {
    case "":
      c.Server.ReplyInstruction(instruction, "")
    case "kill":
      user := c.Server.FindUser(string(instruction.Args[0]))
      c.Server.DeleteUser(c.Server.UserArray[user])
    // TODO: define the sendf feature
    case "end":
      c.Close()
  }
  return nil
}

// Close the server and the client
func (c *ServerClient) Close() {
  for _, user := range c.Server.UserArray {
    c.Server.DeleteUser(user)
  }
  c.Server.Listener.Close()
}
