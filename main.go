package main

import (
	"cli-chat/args"
	"cli-chat/chat"
	"cli-chat/ui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func GetClient(ctx args.CliArgs) chat.Client {
  if ctx.Mode == "client" {
    client, err := chat.OpenConnection(ctx.Ip, ctx.Port, ctx.Name)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
    return client
  }
  server, err := chat.InitServer(ctx.Port)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  client := chat.CreateServerClient(server)
  return client
}

func main() {
  ctx := args.CliArgs{
    Mode: "",
    Ip: "",
    Port: "",
    Name: "",
  }
  args.InitArgs(&ctx)
  args.Parse()
  style := ui.NewStyle("#70EBFF", "#F00057", "#70EBFF")
  data := ui.NewModelData(style)
  client := GetClient(ctx)
  client.Listen()
  program := tea.NewProgram(ui.InitModel(&client, &data))
  // Run the terminal ui
  program.Run()
}
