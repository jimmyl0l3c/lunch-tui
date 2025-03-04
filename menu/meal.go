package menu

import (
	"github.com/charmbracelet/lipgloss"
)

type Meal struct {
	Name string
	Detail string
	Price string
}

var faint = lipgloss.NewStyle().Faint(true)

func (m Meal) String() string {
	return m.Name + "\n" +
		faint.Render(m.Detail) + "\n" +
        m.Price
}
