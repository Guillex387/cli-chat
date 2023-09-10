package chat

// A struct for manage listeners
type Listener struct {
  Stop chan struct{}
  Closed bool
}

// Creates a new Listener
func NewListener() Listener {
  return Listener{Closed: true}
}

// Open a listener
// Return true if was open and false otherwise
func (l *Listener) Open(callback func(stop chan struct{})) bool {
  if !l.Closed {
    return false
  }
  l.Closed = false
  l.Stop = make(chan struct{})
  go callback(l.Stop)
  return true
}

// Send the close signal to the listener
func (l *Listener) Close() {
  if l.Closed {
    return
  }
  l.Stop <- struct{}{}
  close(l.Stop)
  l.Closed = true
}
