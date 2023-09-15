package ins

// Creates a log instruction
func NewLogInstruction(log string) Instruction {
  return NewInstruction("log", []byte(log))
}

// Creates a kill instruction
func NewKillInstruction(userName string) Instruction {
  return NewInstruction("kill", []byte(userName))
}

// Creates a error instruction
func NewErrorInstruction(error string) Instruction {
  return NewInstruction("error", []byte(error))
}

// Creates a end instruction
func NewEndInstruction() Instruction {
  return NewInstruction("end")
}

// Creates a open instruction
func NewOpenInstruction(userName string) Instruction {
  return NewInstruction("open", []byte(userName))
}

// Creates a message instruction
func NewMsgInstruction(userName string, msg string) Instruction {
  return NewInstruction("", []byte(userName), []byte(msg))
}

// Creates a ok instruction
func NewOkInstruction() Instruction {
  return NewInstruction("ok")
}
