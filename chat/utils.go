package chat

import (
	"net"
)

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

// Remove the char '\n' of a string
func RemoveLinebreaks(str string) string {
  result := ""
  for _, char := range str {
    if char != '\n' {
      result += string(char)
    }
  }
  return result
}