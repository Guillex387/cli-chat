package main

import (
	chat "cli-chat/chat"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Constants


var focusStyle = lipgloss.NewStyle().
  Bold(true).
  Foreground(lipgloss.Color("#B48EAD"))

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#BF616A"))

var serverStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#A3BE8C"))

const VIEW_WIDTH = 70

const VIEW_HEIGHT = 40


type TickMsg time.Time

type model struct {
  messages *string
	renderedMessages *bool
  viewPort viewport.Model
  textInput textinput.Model
  err error
}

func initialModel(messages *string, rendered *bool) model {
	textInput := textinput.New()
	textInput.Placeholder = "Write a message/command..."
	textInput.Focus()
	textInput.CharLimit = 300
	textInput.Width = 100
  viewPort := viewport.New(VIEW_WIDTH, VIEW_HEIGHT)
	return model{
    messages: messages,
		textInput: textInput,
    viewPort: viewPort,
		err: nil,
		renderedMessages: rendered,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.CheckMessages(), tea.EnterAltScreen)
}

func (m model) CheckMessages() tea.Cmd {
	return tea.Tick(time.Millisecond * 500, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	switch msg := msg.(type) {
    case tea.KeyMsg:
			m.textInput, tiCmd = m.textInput.Update(msg)
			m.viewPort, vpCmd = m.viewPort.Update(msg)
			switch msg.Type {
				case tea.KeyCtrlC, tea.KeyEsc:
					return m, tea.Quit
				case tea.KeyEnter:
          // TODO: make a parser for the instructions
					// instruction := m.textInput.Value()
					// m.AddMessage(instruction.Id, instruction.Body)
					m.viewPort.SetContent(*m.messages)
					m.textInput.Reset()
					m.viewPort.GotoBottom()
				case tea.KeyUp:
					m.viewPort.YPosition--
				case tea.KeyDown:
					m.viewPort.YPosition++
			}
		case TickMsg:
			if !*m.renderedMessages {
				m.viewPort.SetContent(*m.messages)
				m.viewPort.GotoBottom()
				return m, tea.Batch(tiCmd, vpCmd, m.CheckMessages())
			}
  }
  return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	*m.renderedMessages = true
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewPort.View(),
		m.textInput.View(),
	)
}

// Messages renders

func (m model) AddMessage(sender string, message string) {
	senderWidth := len(sender) + 2
	render := focusStyle.Render(sender + ": ") +
		chat.FormatText(message, VIEW_WIDTH - senderWidth, senderWidth)
	*m.messages += (render + "\n")
	*m.renderedMessages = false
}

func (m model) AddServerMessage(message string) {
	render := serverStyle.Render("Server: ") +
		chat.FormatText(message, VIEW_WIDTH - 8, 8)
	*m.messages += (render + "\n")
	*m.renderedMessages = false
}

func (m model) AddErrorMessage(err string) {
	render := errorStyle.Render(err)
	*m.messages += (render + "\n")
	*m.renderedMessages = false
}

func (m model) AddLogMessage(log string) {
	render := focusStyle.Render(log)
	*m.messages += (render + "\n")
	*m.renderedMessages = false
}

func main() {
	messages := ""
	rendered := false
	model := initialModel(&messages, &rendered)
  program := tea.NewProgram(model, tea.WithAltScreen())
	program.Run()
}
