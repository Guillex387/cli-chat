package ui

type SyntaxError struct {}

func (e SyntaxError) Error() string {
  return "Error parsing the input"
}