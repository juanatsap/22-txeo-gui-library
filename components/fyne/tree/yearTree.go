package tree

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var Months = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
var MonthsEnglish = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var MonthsEnglishMap = map[string]string{
	"Enero":      "January",
	"Febrero":    "February",
	"Marzo":      "March",
	"Abril":      "April",
	"Mayo":       "May",
	"Junio":      "June",
	"Julio":      "July",
	"Agosto":     "August",
	"Septiembre": "September",
	"Octubre":    "October",
	"Noviembre":  "November",
	"Diciembre":  "December",
}
var MonthsSpanishMap = map[string]string{
	"January":   "Enero",
	"February":  "Febrero",
	"March":     "Marzo",
	"April":     "Abril",
	"May":       "Mayo",
	"June":      "Junio",
	"July":      "Julio",
	"August":    "Agosto",
	"September": "Septiembre",
	"October":   "Octubre",
	"November":  "Noviembre",
	"December":  "Diciembre",
}
var Years = []string{"2024", "2023", "2022", "2021"}
var treeData = map[string][]string{
	"":     Years,
	"2021": Months,
	"2022": Months,
	"2023": Months,
	"2024": Months,
}

func MakeTree() *widget.Tree {
	tree := widget.NewTree(
		// Función para obtener los hijos de un nodo
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			if children, ok := treeData[id]; ok {
				// Agregar prefijos al ID de los hijos para mantener la unicidad
				var prefixedChildren []widget.TreeNodeID
				for _, child := range children {
					prefixedChildren = append(prefixedChildren, id+child)
				}
				return prefixedChildren
			}
			return []widget.TreeNodeID{}
		},
		// Función para determinar si un nodo es una rama
		func(id widget.TreeNodeID) bool {
			_, ok := treeData[id]
			return ok
		},
		// Función para crear un nuevo objeto visual para un nodo
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("")
		},
		// Función para actualizar el contenido de un nodo
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			// Extraer el nombre real del nodo
			name := extractName(id)

			// Look if the name is a year
			if name == "2021" || name == "2022" || name == "2023" || name == "2024" || name == "2025" || name == "2026" || name == "2027" {
				label.SetText(name)
			} else {
				// Remove the 4 first characters to get the real name
				label.SetText(name[4:])
			}

			// Center content
			//label.Alignment = fyne.TextAlignCenter
			//label.MinSize() = fyne.NewSize(100, 100)

		})

	return tree
}

func extractName(id string) string {
	//parts := strings.Split(id, "/")
	// return parts[len(parts)-1]
	return id
}
