package main

import (
	"cli-chat/args"
	"cli-chat/chat"
	"cli-chat/ui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Returns the correct client based on the mode (client or server)
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

// Make an initial log in the chat with some data
func InitialLogs(ctx args.CliArgs, data *ui.ModelData) {
  var log string
  if ctx.Mode == "client" {
    log = fmt.Sprintf("Joined to chat at %s:%s", ctx.Ip, ctx.Port)
  } else {
    log = fmt.Sprintf("Waiting connections at %s:%s", chat.GetLocalIp(), ctx.Port)
  }
  data.AddLog(log)
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
  InitialLogs(ctx, &data)
  program := tea.NewProgram(ui.InitModel(&client, &data))
  // Run the terminal ui
  program.Run()
}
