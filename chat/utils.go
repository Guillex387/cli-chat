package chat

import "net"

// Returns the local ip string
func GetLocalIp() string {
  conn, err := net.Dial("udp", "8.8.8.8:80")
  if err != nil {
    return ""
  }
  defer conn.Close()
  localAddr := conn.LocalAddr().(*net.UDPAddr)
  return localAddr.IP.String()
}
