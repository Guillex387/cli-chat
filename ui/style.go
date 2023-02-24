package ui

import "github.com/charmbracelet/lipgloss"

// Represents the style of the ui
type Style struct  {
  FocusColor lipgloss.Style
  ErrorColor lipgloss.Style
  SpecialColor lipgloss.Style
}

// Inits the Style struct
func NewStyle(focusColor string, errorColor string, specialColor string) Style {
  return Style{
    FocusColor: ColorStyle(focusColor),
    ErrorColor: ColorStyle(errorColor),
    SpecialColor: ColorStyle(specialColor),
  }
}

// Converts a color to a lipgloss style
func ColorStyle(color string) lipgloss.Style {
  return lipgloss.NewStyle().Bold(true).
    Foreground(lipgloss.Color(color))
} 

// Render a text with focus style
func (s *Style) RenderFocus(text string) string {
  return s.FocusColor.Render(text)
}

// Render a text with error style
func (s *Style) RenderError(text string) string {
  return s.ErrorColor.Render(text) 
}

// Render a text with special style
func (s *Style) RenderSpecial(text string) string {
  return s.SpecialColor.Render(text)
}
