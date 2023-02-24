package ui

// Represents the data of the model
type ModelData struct {
  Messages string
  RenderedMessages bool
}

// Inits the ModelData struct
func NewModelData() ModelData {
  return ModelData{
    Messages: "",
    RenderedMessages: false,
  }
}

// Adds a message to chat buffer
func (d *ModelData) AddMessage(sender string, message string, style *Style) {
  senderWidth := len(sender) + 2
  buffer := style.RenderFocus(sender + ": ") +
    FormatText(message, VIEW_WIDTH - senderWidth, senderWidth) + "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}

// Adds a error message to chat buffer
func (d *ModelData) AddError(error string, style *Style) {
  buffer := FormatText(style.RenderError(error), VIEW_WIDTH, 0) +
    "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}

// Adds a log message to chat buffer
func (d *ModelData) AddLog(log string, style *Style) {
  buffer := FormatText(style.RenderSpecial(log), VIEW_WIDTH, 0) +
    "\n"
  d.Messages += buffer
  d.RenderedMessages = false
}
