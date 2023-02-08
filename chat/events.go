package chat

// TODO: document this file

type EventListener struct {
  Id int
  callback func(instruction Instruction)
}

type Event struct {
  Seed int
  Listeners []EventListener
}

func (m *Event) CreateListener(callback func(instruction Instruction)) func() {
  id := m.Seed
  m.Listeners = append(m.Listeners, EventListener {Id: id, callback: callback})
  m.Seed++
  return func() {
    m.DeleteListener(id)
  }
}

func (m *Event) DeleteListener(id int) {
  for i, listener := range m.Listeners {
    if listener.Id == id {
      m.Listeners = append(m.Listeners[0:i], m.Listeners[i+1:]...)
    }
  }
}

func (m *Event) Trigger(instruction Instruction) {
  for _, listener := range m.Listeners {
    listener.callback(instruction)
  }
}
