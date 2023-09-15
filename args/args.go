package args

import "flag"

const DESCRIPTION = "Terminal based chat with multiple functionality" +
                    " which has its own data transfer protocol.\n"
const MODE_HELP = "Choose the execution mode of the program, \"client\", \"server\""

const PORT_HELP = "The port to listen or connect to a server"

const IP_HELP = "IP direction to connect to a host (client only)"

const NAME_HELP = "The display name in the chat (client only)"

type CliArgs struct {
  Mode string
  Ip string
  Port string
  Name string
}

func InitArgs(cli *CliArgs) {
  flag.StringVar(&cli.Mode, "mode", "client", MODE_HELP)
  flag.StringVar(&cli.Port, "port", "8000", PORT_HELP)
  flag.StringVar(&cli.Ip, "ip", "127.0.0.1", IP_HELP)
  flag.StringVar(&cli.Name, "name", "unknown", NAME_HELP)

  flag.Usage = func() {
    println(DESCRIPTION)

    println("Options:")
    flag.PrintDefaults()
  }
}

func Parse() {
  if flag.Parsed() {
    return
  }

  flag.Parse()
}
