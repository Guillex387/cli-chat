package ui

type SyntaxError struct {}

func (e SyntaxError) Error() string {
  return "Syntax error in the input instruction"
}
