package menu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

type Meal struct {
	Name   string
	Detail string
	Price  string
}

func (m Meal) String(maxWidth int) string {
	return fmt.Sprintf("%s\n%s\n%s",
		styles.HighlightStyle.Width(maxWidth).Render(m.Name),
		styles.FaintStyle.Width(maxWidth).Render(m.Detail),
		m.Price)
}

func Menu(meals []Meal, maxWidth int) string {
	l := list.New().
		Enumerator(func(_ list.Items, i int) string {
			return ""
		}).
		ItemStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			return lipgloss.NewStyle().MarginBottom(1)
		}).
		EnumeratorStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			return lipgloss.NewStyle()
		})

	for _, d := range meals {
		l.Item(d.String(maxWidth))
	}

	return l.String()
}
