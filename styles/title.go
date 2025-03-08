package styles

import "github.com/charmbracelet/lipgloss"

var (
	Divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	Url = lipgloss.NewStyle().Foreground(special).Render

	Title = BaseStyle.MarginTop(1).Render

	Subtitle = BaseStyle.
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle).
			Render
)
