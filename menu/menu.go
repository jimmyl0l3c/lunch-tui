package menu

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var faint = lipgloss.NewStyle().Faint(true)

var hightlightColor = lipgloss.Color("#EE6FF8")

var mealHighlight = lipgloss.NewStyle().Foreground(hightlightColor)

type Meal struct {
	Name   string
	Detail string
	Price  string
}

func (m Meal) String(maxWidth int) string {
	return mealHighlight.Width(maxWidth).Render(m.Name) + "\n" +
		faint.Width(maxWidth).Render(m.Detail) + "\n" +
		m.Price
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
