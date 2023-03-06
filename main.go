package main

// Example of client for the ui

import (
	"cli-chat/chat"
	"cli-chat/ui"
	"fmt"

	// tea "github.com/charmbracelet/bubbletea"
)

type ClientEx struct {
  SendEvent chat.Event
}

func NewClient() chat.Client {
  return &ClientEx{SendEvent: chat.NewEvent()}
}

func (c *ClientEx) Event() *chat.Event {
  return &c.SendEvent
}

func (c *ClientEx) SendInstruction(instruction chat.Instruction) error {
  c.SendEvent.Trigger(instruction)
  return nil
}

func (c *ClientEx) Close() {
  c.SendEvent.Clear()
}

func main() {
  mode := ""
  ip := ""
  port := ""
  ui.InitArgs(&mode, &ip, &port)
  ui.Parse()
  fmt.Printf("%s -> %s:%s\n", mode, ip, port)
  // client := NewClient()
  // style := ui.NewStyle("#70EBFF", "#F00057", "#70EBFF")
  // data := ui.NewModelData(style)
  // program := tea.NewProgram(ui.InitModel(&client, &data))
  // client.Event().On("", func(this chat.EventListener, instruction chat.Instruction) {
  //   data.AddMessage("You", string(instruction.Args[1]))
  // })
  // client.Event().On("error", func(this chat.EventListener, instruction chat.Instruction) {
  //   data.AddError(string(instruction.Args[0]))
  // })
  // client.Event().On("end", func(this chat.EventListener, instruction chat.Instruction) {
  //   program.Quit()
  // })
  // program.Run()
  // TODO: define the main with an argument parser
}
