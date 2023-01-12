package main

import (
	chat "cli-chat/chat"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var focusStyle = lipgloss.NewStyle().
  Bold(true).
  Foreground(lipgloss.Color("#7D56F4"))

func main() {
  program := tea.NewProgram(initialModel())
	program.Run()
}

type model struct {
  messages []string
  viewPort viewport.Model
  textInput textinput.Model
  err error
}

func initialModel() model {
	input := textinput.New()
	input.Placeholder = "Write a message..."
	input.Focus()
	input.CharLimit = 300
	input.Width = 100
  view := viewport.New(65, 40)
	return model{
    messages: make([]string, 0),
		textInput: input,
    viewPort: view,
		err: nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textInput, tiCmd = m.textInput.Update(msg)
	m.viewPort, vpCmd = m.viewPort.Update(msg)

	switch msg := msg.(type) {
    case tea.KeyMsg:
    switch msg.Type {
      case tea.KeyCtrlC, tea.KeyEsc:
        return m, tea.Quit
      case tea.KeyEnter:
        m.messages = append(m.messages, focusStyle.Render("You: ") + chat.FormatText(m.textInput.Value(), 60, 5))
        m.viewPort.SetContent(strings.Join(m.messages, "\n"))
        m.textInput.Reset()
        m.viewPort.GotoBottom()
    }
    return m, tea.Batch(tiCmd, vpCmd)
  }
  return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewPort.View(),
		m.textInput.View(),
	)
}