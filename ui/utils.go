package ui

// A function to add line breaks
func FormatText(str string, breakPos int) string {
  result := ""
  counter := 0
  for _, char := range str {
    result += string(char)
    if char == '\n' {
      counter = 0
    }
    if counter == (breakPos - 1) {
      result += "\n"
      counter = 0
    }
    counter++
  }
  return result
}
