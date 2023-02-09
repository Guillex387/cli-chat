package chat

// Represent a listener for a event
type EventListener struct {
  Id int
  callback func(instruction Instruction)
}

// Represent a event 
type Event struct {
  Seed int
  Listeners []EventListener
}

// Creates a new event
func NewEvent() Event {
  return Event {0, make([]EventListener, 0)}
}

// Creates a listener for the event and return a deleter
func (m *Event) CreateListener(callback func(instruction Instruction)) func() {
  id := m.Seed
  m.Listeners = append(m.Listeners, EventListener {Id: id, callback: callback})
  m.Seed++
  return func() {
    m.DeleteListener(id)
  }
}

// Deletes a listener from the event
func (m *Event) DeleteListener(id int) {
  for i, listener := range m.Listeners {
    if listener.Id == id {
      m.Listeners = append(m.Listeners[0:i], m.Listeners[i+1:]...)
    }
  }
}

// Activate the event and executes all the listeners 
func (m *Event) Trigger(instruction Instruction) {
  for _, listener := range m.Listeners {
    go listener.callback(instruction)
  }
}

// Deletes all the listener of the Event
func (m *Event) Clear() {
  m.Listeners = make([]EventListener, 0)
}
