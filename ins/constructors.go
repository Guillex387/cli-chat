package ins

// Creates a log instruction
func NewlogInstruction(log string) Instruction {
  return NewIntruction("log", []byte(log))
}

// Creates a kill instruction
func NewKillInstruction(userName string) Instruction {
  return NewIntruction("kill", []byte(userName))
}

// Creates a error instruction
func NewErrorInstruction(error string) Instruction {
  return NewIntruction("error", []byte(error))
}

// Creates a end instruction
func NewEndInstruction() Instruction {
  return NewIntruction("end")
}

// Creates a open instruction
func NewOpenInstruction(userName string) Instruction {
  return NewIntruction("open", []byte(userName))
}

// Creates a message instruction
func NewMsgInstruction(userName string, msg string) Instruction {
  return NewIntruction("", []byte(userName), []byte(msg))
}

// Creates a ok instruction
func NewOkInstruction() Instruction {
  return NewIntruction("ok")
}
