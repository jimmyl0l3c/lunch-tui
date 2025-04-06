package menu

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

type PrintableColumn interface {
	Render(maxWidth int) string
}

type RestaurantData struct {
	Name  string
	Meals []Meal
}

func (restaurant RestaurantData) Render(maxWidth int) string {
	content := Menu(restaurant.Meals, maxWidth)

	return lipgloss.JoinVertical(lipgloss.Left, styles.ListHeader(restaurant.Name), content)
}

type RestaurantError struct {
	RestaurantName string
	Msg            string
}

func (e *RestaurantError) Render(maxWidth int) string {
	content := styles.ErrorStyle.Width(maxWidth).Render("Error:", e.Msg)

	return lipgloss.JoinVertical(lipgloss.Left, styles.ListHeader(e.RestaurantName), content)
}

func RestaurantRow(restaurants []PrintableColumn, physicalWidth int) (row string) {
	maxMenuWidth := (physicalWidth - 4 - (6 * len(restaurants))) / len(restaurants)

	lastIndex := len(restaurants) - 1

	for i, restaurant := range restaurants {
		column := restaurant.Render(maxMenuWidth)

		if i == 0 {
			row = styles.List(column)
		} else if i == lastIndex {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, column)
		} else {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, styles.List(column))
		}
	}

	return
}
