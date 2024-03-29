package ui

import (
	chat "cli-chat/chat"
  "cli-chat/ins"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Viewport constants
const VIEW_WIDTH = 70
const VIEW_HEIGHT = 25

// Time interval to check the new data
const REFRESH_TIME = 50 * time.Millisecond

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
  textInput.CharLimit = VIEW_WIDTH
  textInput.Width = VIEW_WIDTH
  viewPort := viewport.New(VIEW_WIDTH, VIEW_HEIGHT)

  // Setup the client listeners
  (*client).Event().On("", func(this chat.EventListener, instruction ins.Instruction) {
    data.AddMessage(string(instruction.Args[0]), string(instruction.Args[1]))
  })
  (*client).Event().On("log", func(this chat.EventListener, instruction ins.Instruction) {
    data.AddLog(string(instruction.Args[0]))
  })
  (*client).Event().On("error", func(this chat.EventListener, instruction ins.Instruction) {
    data.AddError(string(instruction.Args[0]))
  })
  (*client).Event().On("clear", func(this chat.EventListener, instruction ins.Instruction) {
    data.Clear()
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
  return tea.Tick(REFRESH_TIME, func(t time.Time) tea.Msg {
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
          (*m.Client).SendInstruction(ins.NewEndInstruction())
          return m, tea.Quit
        
        case tea.KeyEnter:
          input := m.TextInput.Value()
          if input == "" {
            break
          }
          if m.OnInstructionInput(input) == "end" {
            return m, tea.Quit
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

// Executes when the user put a input instruction
// and returns the id
func (m Model) OnInstructionInput(input string) string {
  // Parse the input to a instruction
  instruction, err := ins.ParseInstruction(input).ToInstruction()
  if err != nil {
    (*m.Data).AddError(err.Error())
  } else {
    // Reply the instruction to the client
    (*m.Client).SendInstruction(instruction)
  }
  return instruction.Id
}

// Render the model
func (m Model) View() string {
  return fmt.Sprintf(
    "%s\n\n%s",
    m.ViewPort.View(),
    m.TextInput.View(),
  )
}
