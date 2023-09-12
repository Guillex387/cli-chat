package chat

import "cli-chat/ins"

// Represent a listener for a event
type EventListener struct {
  Id int
  InstructionId []string
  Callback func(this EventListener, instruction ins.Instruction)
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

// Creates a listener of a certain instruction
func (m *Event) On(instructionId string, callback func(this EventListener, instruction ins.Instruction)) EventListener {
  listener := EventListener{Id: m.Seed, InstructionId: []string{instructionId}, Callback: callback}
  m.Listeners = append(m.Listeners, listener)
  m.Seed++
  return listener
}

// Creates a listener of a multiple instructions types
func (m *Event) OnMultiple(instructionsIds []string, callback func(this EventListener, instruction ins.Instruction)) {
  listener := EventListener{Id: m.Seed, InstructionId: instructionsIds, Callback: callback}
  m.Listeners = append(m.Listeners, listener)
  m.Seed++
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
func (m *Event) Trigger(instruction ins.Instruction) {
  for _, listener := range m.Listeners {
    for _, id := range listener.InstructionId {
      if instruction.Id == id {
        go listener.Callback(listener, instruction)
        break
      }
    }
  }
}

// Deletes all the listener of the Event
func (m *Event) Clear() {
  m.Listeners = make([]EventListener, 0)
}
