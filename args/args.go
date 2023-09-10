package args

import "flag"

type CliArgs struct {
  Mode string
  Ip string
  Port string
  Name string
}

func InitArgs(cli *CliArgs) {
  flag.StringVar(&cli.Mode, "mode", "client", "Value: \"client\", \"server\". Choose the execution mode of the program")
  flag.StringVar(&cli.Port, "port", "8000", "The port to listen or connect to a server")
  flag.StringVar(&cli.Ip, "ip", "", "IP direction to connect to a host (client only)")
  flag.StringVar(&cli.Name, "name", "unknow", "The display name in the chat (client only)")
}

func Parse() {
  if !flag.Parsed() {
    flag.Parse()
  }
}
