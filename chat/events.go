package chat

// Represent a listener for a event
type EventListener struct {
  Id int
  InstructionId string
  Callback func(this EventListener, instruction Instruction)
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

// Creates a listeners of a certain instruction
func (m *Event) On(instructionId string, callback func(this EventListener, instruction Instruction)) EventListener {
  listener := EventListener{Id: m.Seed, InstructionId: instructionId, Callback: callback}
  m.Listeners = append(m.Listeners, listener)
  m.Seed++
  return listener
}

// Deletes a listener from the event
func (m *Event) DeleteListener(listener EventListener) {
  for i, listenerElement := range m.Listeners {
    if listenerElement.Id == listener.Id {
      m.Listeners = append(m.Listeners[0:i], m.Listeners[i+1:]...)
    }
  }
}

// Activate the event and executes all the listeners 
func (m *Event) Trigger(instruction Instruction) {
  for _, listener := range m.Listeners {
    if instruction.Id == listener.InstructionId {
      go listener.Callback(listener, instruction)
    }
  }
}

// Deletes all the listener of the Event
func (m *Event) Clear() {
  m.Listeners = make([]EventListener, 0)
}
