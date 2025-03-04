package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
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

var meals = []Meal {
    {"Gulášová polévka s hlívou", "Gulášová polévka s hlívou, paprikou a kukuřicí", "50 Kč"},
    {"1: Zeleninové kari s batáty. Basmati rýže", "Svěží kari s batáty, paprikou, fawa fazolemi a ananasem. Basmati rýže. Marinované tofu.", "165 Kč"},
    {"2: Řecké ragú s lilky. Bulgur. Feta.", "Řecké ragú se sojovými výpečky, lilky, cuketou, rajčaty a olivami. Feta sýr. Bulgur.", "165 Kč"},
    {"3: Salát z červené řepy. Sýr. Pečivo.", "Salát z červené řepy, fenyklu a pomerančů. Kozí sýr. Naan.", "165 Kč"},
}

const selected = 1

func main() {
	baseStyle := lipgloss.NewStyle().
		MarginBottom(1).
		MarginLeft(1)
	dimColor := lipgloss.Color("250")
	hightlightColor := lipgloss.Color("#EE6FF8")

	l := list.New().
		Enumerator(func(_ list.Items, i int) string {
			if i == selected {
				return "│\n│\n│"
			}
			return " "
		}).
		ItemStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			st := baseStyle
			if selected == i {
				return st.Foreground(hightlightColor)
			}
			return st.Foreground(dimColor)
		}).
		EnumeratorStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			if selected == i {
				return lipgloss.NewStyle().Foreground(hightlightColor)
			}
			return lipgloss.NewStyle().Foreground(dimColor)
		})

	for _, d := range meals {
		l.Item(d.String())
	}

	fmt.Println()
	fmt.Println(l)
}
