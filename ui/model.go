package ui

import (
	chat "cli-chat/chat"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Viewport constants
const VIEW_WIDTH = 70
const VIEW_HEIGHT = 40

// Represent a event of type tick
type TickMsg time.Time

// TODO: implent it
// Represents the style of the ui
type Style struct  {
  FocusColor string
  ErrorColor string
  SpecialColor string
}

// Represents the data of the model
type ModelData struct {
  Messages string
  RenderedMessages bool
}

// Represents the ui model
type Model struct {
  Data *ModelData
  ViewPort viewport.Model
  TextInput textinput.Model
}

// Inits the model
func InitModel(client chat.Client, data *ModelData) Model {
  textInput := textinput.New()
  textInput.Placeholder = "Write a message/command..."
  textInput.Focus()
  textInput.CharLimit = 300
  textInput.Width = 100
  viewPort := viewport.New(VIEW_WIDTH, VIEW_HEIGHT)
  return Model{
    Data: data,
    TextInput: textInput,
    ViewPort: viewPort,
  }
}

// Inits the model loop
func (m Model) Init() tea.Cmd {
  return tea.Batch(textinput.Blink, m.CheckMessages(), tea.EnterAltScreen)
}

// Check if are not rendered messages
func (m Model) CheckMessages() tea.Cmd {
  return tea.Tick(time.Millisecond * 500, func(t time.Time) tea.Msg {
    return TickMsg(t)
  })
}

// Callback of ui update
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var (
    textInputCmd tea.Cmd
    viewPortCmd tea.Cmd
  )
  // Event switch
  switch msg := msg.(type) {
    case tea.KeyMsg:
      m.TextInput, textInputCmd = m.TextInput.Update(msg)
      m.ViewPort, viewPortCmd = m.ViewPort.Update(msg)
      // Control switch
      switch msg.Type {
        case tea.KeyCtrlC, tea.KeyEsc:
          return m, tea.Quit
        case tea.KeyEnter:
          m.ViewPort.SetContent(m.Data.Messages)
          m.TextInput.Reset()
          m.ViewPort.GotoBottom()
        case tea.KeyUp:
          m.ViewPort.YPosition--
        case tea.KeyDown:
          m.ViewPort.YPosition++
      }
    case TickMsg:
      if !m.Data.RenderedMessages {
        m.ViewPort.SetContent(m.Data.Messages)
        m.ViewPort.GotoBottom()
        return m, tea.Batch(textInputCmd, viewPortCmd, m.CheckMessages())
      }
  }
  return m, tea.Batch(textInputCmd, viewPortCmd)
}

// The render of the modell into a string
func (m Model) View() string {
  m.Data.RenderedMessages = true
  return fmt.Sprintf(
    "%s\n\n%s",
    m.ViewPort.View(),
    m.TextInput.View(),
  )
}
