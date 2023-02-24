package chat

type OpenConnectionError struct {}

func (e *OpenConnectionError) Error() string {
  return "Error opening the connection"
}

type OpenServerError struct {}

func (e *OpenServerError) Error() string {
  return "Error creating the server"
}

type ConnectionIOError struct {}

func (e *ConnectionIOError) Error() string {
  return "Error reading or writing in the connection"
}
