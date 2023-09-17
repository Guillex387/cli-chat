package chat

import "cli-chat/ins"

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
func (c *ServerClient) SendInstruction(instruction ins.Instruction) {
  c.Server.ManageServerInstruction(instruction)
}

// Listen to new user connections
func (c *ServerClient) Listen() {
  c.Server.Listen()
}

// Close the server and the client
func (c *ServerClient) Close() {
  c.Server.Close()
}
