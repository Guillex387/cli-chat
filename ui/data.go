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

// Adds a message to chat buffer
func (d *ModelData) AddMessage(sender string, message string) {
  senderWidth := len(sender) + 2
  buffer := d.Style.RenderFocus(sender + ": ") +
    FormatText(message, VIEW_WIDTH - senderWidth, senderWidth) + "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}

// Adds a error message to chat buffer
func (d *ModelData) AddError(error string) {
  buffer := FormatText(d.Style.RenderError(error), VIEW_WIDTH, 0) +
    "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}

// Adds a log message to chat buffer
func (d *ModelData) AddLog(log string) {
  buffer := FormatText(d.Style.RenderSpecial(log), VIEW_WIDTH, 0) +
    "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}
