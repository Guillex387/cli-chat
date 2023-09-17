package ui

// Represents the data of the model
type ModelData struct {
  Style Style
  Messages string
  RenderedMessages bool
}

// Inits the ModelData struct
func NewModelData(style Style) ModelData {
  return ModelData{
    Messages: "",
    RenderedMessages: false,
    Style: style,
  }
}

// Adds string data to chat buffer
func (d *ModelData) AddData(buffer string) {
  buffer += "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}

// Clear the message buffer
func (d *ModelData) Clear() {
  d.Messages = ""
  d.RenderedMessages = false
}

// Adds a message to chat buffer
func (d *ModelData) AddMessage(sender string, message string) {
  senderRendered := d.Style.RenderFocus(sender + ": ")
  buffer := FormatText(senderRendered + message, VIEW_WIDTH)
  d.AddData(buffer)
}

// Adds a error message to chat buffer
func (d *ModelData) AddError(error string) {
  errorRendered := d.Style.RenderError(error)
  buffer := FormatText(errorRendered, VIEW_WIDTH)
  d.AddData(buffer)
}

// Adds a log message to chat buffer
func (d *ModelData) AddLog(log string) {
  buffer := d.Style.RenderSpecial(log)
  d.AddData(buffer)
}
