package styles

import "github.com/charmbracelet/lipgloss"

var (
	List = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		PaddingRight(2).
		MarginRight(2).
		Height(8).
		Render

	ListHeader = BaseStyle.
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render
)
