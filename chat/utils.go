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

// A function to add line breaks
func FormatText(str string, breakPos int, margin int) string {
  result := ""
  separator := ""
  separator += "\n"
  for i := 0; i < margin; i++ {
    separator += " "
  }
  for i, char := range str {
    result += string(char)
    if i == (breakPos - 1) {
      result += separator
    }
  }
  return result
}