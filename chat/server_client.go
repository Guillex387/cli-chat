package chat

import "fmt"

// Struct that represent the server client
type ServerClient struct {
  Server Server
}

// Creates a client for a server
func CreateServerClient(server Server) Client {
  return &ServerClient {Server: server}
}

// Getter of the event manager
func (c *ServerClient) Event() *Event {
  return &c.Server.SendEvent
}

// Send a instruction to server
func (c *ServerClient) SendInstruction(instruction Instruction) error {
  switch instruction.Id {
    case "":
      instruction.Args[0] = []byte("Server")
      c.Server.ReplyInstruction(instruction, "")
    case "kill":
      user := c.Server.FindUser(string(instruction.Args[0]))
      if user == -1 {
        errorMsg := fmt.Sprintf("User '%s' not found", string(instruction.Args[0]))
        c.Server.SendEvent.Trigger(NewErrorInstruction(errorMsg))
        break
      }
      c.Server.DeleteUser(c.Server.UserArray[user])
    case "end":
      c.Server.SendEvent.Trigger(instruction)
      c.Close()
    // TODO: define the sendf feature
  }
  return nil
}

// Listen to new user connections
func (c *ServerClient) Listen() {
  c.Server.Listen()
}

// Close the server and the client
func (c *ServerClient) Close() {
  c.Server.DeleteAllUsers()
  c.Server.SendEvent.Clear()
  c.Server.ConnListener.Close()
  c.Server.Listener.Close()
}
