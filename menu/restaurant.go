package menu

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

type Restaurant struct {
	Name  string
	Meals []Meal
	Err   error
}

func RestaurantColumn(restaurant Restaurant, maxWidth int) string {
	var content string

	if restaurant.Err == nil {
		content = Menu(restaurant.Meals, maxWidth)
	} else {
		content = styles.ErrorStyle.Width(maxWidth).Render("Error:", restaurant.Err.Error())
	}

	return lipgloss.JoinVertical(lipgloss.Left, styles.ListHeader(restaurant.Name), content)
}

func RestaurantRow(restaurants []Restaurant, physicalWidth int) (row string) {
	maxMenuWidth := (physicalWidth - 4 - (6 * len(restaurants))) / len(restaurants)

	lastIndex := len(restaurants) - 1

	for i, restaurant := range restaurants {
		column := RestaurantColumn(restaurant, maxMenuWidth)

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
