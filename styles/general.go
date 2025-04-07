package styles

import "github.com/charmbracelet/lipgloss"

var (
	normal     = lipgloss.Color("#EEEEEE")
	subtle     = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	hightlight = lipgloss.Color("#F371FC")
	special    = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	error      = lipgloss.Color("#EB6F92")
	faint      = lipgloss.Color("#A8A8A8")

	BaseStyle      = lipgloss.NewStyle().Foreground(normal)
	HighlightStyle = lipgloss.NewStyle().Foreground(hightlight)
	DetailStyle    = lipgloss.NewStyle().Foreground(faint)

	ErrorStyle = lipgloss.NewStyle().Foreground(error)

	Error = ErrorStyle.Render
)
