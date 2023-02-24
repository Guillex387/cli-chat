package ui

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
