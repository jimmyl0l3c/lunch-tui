package menu

import "github.com/charmbracelet/lipgloss"

type Restaurant struct {
	Name  string
	Meals []Meal
}

var (
	listStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(subtle).
			PaddingRight(2).
			MarginRight(2).
			Height(8)

	listHeader = base.
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render
)

func RestaurantColumn(restaurant Restaurant, maxWidth int) string {
	return lipgloss.JoinVertical(lipgloss.Left,
		listHeader(restaurant.Name),
		Menu(restaurant.Meals, maxWidth),
	)
}

func RestaurantRow(restaurants []Restaurant, physicalWidth int) string {
	var maxMenuWidth = (physicalWidth - 4 - (6 * len(restaurants))) / len(restaurants)

	var lastIndex = len(restaurants) - 1

	var row = ""

	for i, restaurant := range restaurants {
		column := RestaurantColumn(restaurant, maxMenuWidth)

		if i == 0 {
			row = listStyle.Render(column)
		} else if i == lastIndex {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, column)
		} else {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, listStyle.Render(column))
		}
	}

	return row
}
