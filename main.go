package main

import (
	"fmt"
	"os"
	"strings"

    "lunch/menu"
    "github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var meals = []menu.Meal {
    {"Gulášová polévka s hlívou", "Gulášová polévka s hlívou, paprikou a kukuřicí", "50 Kč"},
    {"1: Zeleninové kari s batáty. Basmati rýže s velmi dlouhým názvem", "Svěží kari s batáty, paprikou, fawa fazolemi a ananasem. Basmati rýže. Marinované tofu.", "165 Kč"},
    {"2: Řecké ragú s lilky. Bulgur. Feta.", "Řecké ragú se sojovými výpečky, lilky, cuketou, rajčaty a olivami. Feta sýr. Bulgur.", "165 Kč"},
    {"3: Salát z červené řepy. Sýr. Pečivo.", "Salát z červené řepy, fenyklu a pomerančů. Kozí sýr. Naan.", "165 Kč"},
}

var restaurants = []menu.Restaurant {
    {"Restaurace 1", meals},
    {"Restaurace 2", meals},
    {"Restaurace 3", meals},
}

// Style definitions.
var (
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func main() {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
    scraperVersion := "v0.1.0"

    menu.PrintTitle(scraperVersion, physicalWidth)

	doc := strings.Builder{}

    doc.WriteString(menu.RestaurantRow(restaurants, physicalWidth))

	fmt.Println(docStyle.Render(doc.String()))
}
