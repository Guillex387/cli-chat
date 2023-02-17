package ui

import (
	chat "cli-chat/chat"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: document this file and find the solution for
//       the pointer bug in the interface tea.Model

const VIEW_WIDTH = 70
const VIEW_HEIGHT = 40

type TickMsg time.Time

type Style struct  {
  FocusColor string
  ErrorColor string
  SpecialColor string
}

type Model struct {
  Messages string
  RenderedMessages bool
  ViewPort viewport.Model
  TextInput textinput.Model
  Error error
}

func InitModel(client chat.Client) Model {
  textInput := textinput.New()
  textInput.Placeholder = "Write a message/command..."
  textInput.Focus()
  textInput.CharLimit = 300
  textInput.Width = 100
  viewPort := viewport.New(VIEW_WIDTH, VIEW_HEIGHT)
  return Model{
    Messages: "",
    RenderedMessages: false,
    TextInput: textInput,
    ViewPort: viewPort,
    Error: nil,
  }
}

func (m* Model) Init() tea.Cmd {
  return tea.Batch(textinput.Blink, m.CheckMessages(), tea.EnterAltScreen)
}

func (m* Model) CheckMessages() tea.Cmd {
  return tea.Tick(time.Millisecond * 500, func(t time.Time) tea.Msg {
    return TickMsg(t)
  })
}

func (m* Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
          return *m, tea.Quit
        case tea.KeyEnter:
          m.ViewPort.SetContent(m.Messages)
          m.TextInput.Reset()
          m.ViewPort.GotoBottom()
        case tea.KeyUp:
          m.ViewPort.YPosition--
        case tea.KeyDown:
          m.ViewPort.YPosition++
      }
    case TickMsg:
      if !m.RenderedMessages {
        m.ViewPort.SetContent(m.Messages)
        m.ViewPort.GotoBottom()
        return *m, tea.Batch(textInputCmd, viewPortCmd, m.CheckMessages())
      }
  }
  return *m, tea.Batch(textInputCmd, viewPortCmd)
}

func (m* Model) View() string {
  m.RenderedMessages = true
  return fmt.Sprintf(
    "%s\n\n%s",
    m.ViewPort.View(),
    m.TextInput.View(),
  )
}
