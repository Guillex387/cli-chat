package ins

type SyntaxError struct {}

func (e SyntaxError) Error() string {
  return "Syntax error in the input instruction"
}
