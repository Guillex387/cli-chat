package ui

import "flag"

func InitArgs(mode *string, ip *string, port *string) {
  flag.StringVar(mode, "mode", "client", "Value: \"client\", \"server\" (Default client). Choose the execution mode of the program")
  flag.StringVar(ip, "ip", "", "IP direction to connect to a host (client only)")
  flag.StringVar(port, "port", "8000", "The port to listen or connect to a server (Default 8000)")
}

func Parse() {
  if !flag.Parsed() {
    flag.Parse()
  }
}
