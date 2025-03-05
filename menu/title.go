package menu

import (
    "fmt"

    "github.com/charmbracelet/lipgloss"
)

// Style definitions.
var (
    url = lipgloss.NewStyle().Foreground(special).Render

    divider = lipgloss.NewStyle().
        SetString("•").
        Padding(0, 1).
        Foreground(subtle).
        String()

    titleStyle = base.MarginTop(1)

    sourceStyle = base.
        BorderStyle(lipgloss.NormalBorder()).
        BorderTop(true).
        BorderForeground(subtle)
)

func PrintTitle(version string, physicalWidth int) {
    title := lipgloss.JoinVertical(lipgloss.Center,
        titleStyle.Render("Edhouse lunch menu scraper"+divider+url(version)),
        sourceStyle.Render("Source available"+divider+url("https://github.com/jimmyl0l3c/lunch-tui")),
    )

    centeredTitle := lipgloss.PlaceHorizontal(physicalWidth, lipgloss.Center, title)

    fmt.Println(centeredTitle+"\n")
}
