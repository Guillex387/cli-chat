package ui

import "flag"

func InitArgs(mode *string, ip *string, port *string) {
  flag.StringVar(mode, "mode", "client", "mode: \"client\", \"server\" (Default client). Choose the mode of execution of the program")
  flag.StringVar(ip, "ip", "", "ip: pass the ip to connect to a host (client only)")
  flag.StringVar(port, "port", "8000", "port: the port to listen or connect to a server (Default 8000)")
}

func Parse() {
  if !flag.Parsed() {
    flag.Parse()
  }
}
