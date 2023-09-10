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
const ViewWidth = 70
const ViewHeight = 40

// Represent a event of type tick
type TickMsg time.Time

// Represents the ui model
type Model struct {
  Client *chat.Client
  Data *ModelData
  ViewPort viewport.Model
  TextInput textinput.Model
}

// Inits the model
func InitModel(client *chat.Client, data *ModelData) Model {
  // Components config
  textInput := textinput.New()
  textInput.Placeholder = "Write a message/command..."
  textInput.Focus()
  textInput.CharLimit = 300
  textInput.Width = 100
  viewPort := viewport.New(ViewWidth, ViewHeight)

  // Setup the client listeners
  (*client).Event().On("", func(this chat.EventListener, instruction chat.Instruction) {
    data.AddMessage(string(instruction.Args[0]), string(instruction.Args[1]))
  })
  (*client).Event().On("log", func(this chat.EventListener, instruction chat.Instruction) {
    data.AddLog(string(instruction.Args[0]))
  })
  (*client).Event().On("error", func(this chat.EventListener, instruction chat.Instruction) {
    data.AddError(string(instruction.Args[0]))
  })

  return Model{
    Client: client,
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
  return tea.Tick(time.Millisecond * 50, func(t time.Time) tea.Msg {
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
          (*m.Client).SendInstruction(chat.NewEndInstruction())
          return m, tea.Quit
        
        case tea.KeyEnter:
          text := m.TextInput.Value()
          // Parse the input to a instruction
          instruction, err := ParseInstruction(text).ToInstruction()
          if err != nil {
            (*m.Data).AddError(err.Error())
          } else {
            // Reply the instruction to the client
            (*m.Client).SendInstruction(instruction)
            if instruction.Id == "end" {
              return m, tea.Quit
            }
          }
          m.TextInput.Reset()
          m.ViewPort.GotoBottom()
        
        case tea.KeyUp:
          m.ViewPort.YPosition--
        
        case tea.KeyDown:
          m.ViewPort.YPosition++
      }
    
    case TickMsg:
      if !m.Data.RenderedMessages {
        m.Data.RenderedMessages = true
        m.ViewPort.SetContent(m.Data.Messages)
        m.ViewPort.GotoBottom()
        return m, tea.Batch(textInputCmd, viewPortCmd, m.CheckMessages())
      }
  }
  return m, tea.Batch(textInputCmd, viewPortCmd, m.CheckMessages())
}

// Render the model
func (m Model) View() string {
  return fmt.Sprintf(
    "%s\n\n%s",
    m.ViewPort.View(),
    m.TextInput.View(),
  )
}
