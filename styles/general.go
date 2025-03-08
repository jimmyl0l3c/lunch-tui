package styles

import "github.com/charmbracelet/lipgloss"

var (
	normal     = lipgloss.Color("#EEEEEE")
	subtle     = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	hightlight = lipgloss.Color("#EE6FF8")
	special    = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	error      = lipgloss.Color("#EB6F92")

	BaseStyle      = lipgloss.NewStyle().Foreground(normal)
	HighlightStyle = lipgloss.NewStyle().Foreground(hightlight)
	FaintStyle     = lipgloss.NewStyle().Faint(true)

	Error = lipgloss.NewStyle().Foreground(error).Render
)
