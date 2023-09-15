package ins

// Remove the char '\n' of a string
func RemoveLineBreaks(str string) string {
  result := ""
  for _, char := range str {
    if char != '\n' {
      result += string(char)
    }
  }
  return result
}
