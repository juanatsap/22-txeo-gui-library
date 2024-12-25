package calendar

import (
	"fmt"
	"image/color"
	"time"
	"txeo-gui-library/models"
	"txeo-gui-library/styles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

// MakeCustomCalendarII creates a fully custom calendar
func MakeCustomCalendar(selectedDate time.Time, blocks models.Blocks) *fyne.Container {

	startingDate := time.Now()
	// Title for the month and year
	monthLabel := widget.NewLabelWithStyle(
		startingDate.Format("January 2006"),
		fyne.TextAlignCenter, fyne.TextStyle{Bold: true},
	)
	monthLabel.Alignment = fyne.TextAlignCenter

	// Days of the week (Monday to Sunday)
	weekdays := []string{"L", "M", "M", "J", "V", "S", "D"}
	weekdayLabels := make([]fyne.CanvasObject, len(weekdays))
	for i, day := range weekdays {
		label := canvas.NewText(day, color.Black)
		label.Alignment = fyne.TextAlignCenter
		weekdayLabels[i] = label
	}

	// Container for days of the week
	weekdayRow := container.NewGridWithColumns(7, weekdayLabels...)

	// Generate calendar cells
	cells := createCustomCalendarCells(startingDate, blocks)

	// Main calendar grid
	calendarGrid := container.NewGridWithColumns(7, cells...)

	// Combine everything into a vertical box
	calendarContainer := container.NewVBox(
		monthLabel,   // Month and year
		weekdayRow,   // Days of the week
		calendarGrid, // Calendar days
	)

	return calendarContainer
}

// createCustomCalendarCells generates the cells for the custom calendar
func createCustomCalendarCells(startingDate time.Time, blocks models.Blocks) []fyne.CanvasObject {
	today := time.Now() // Current day
	firstDayOfMonth := time.Date(startingDate.Year(), startingDate.Month(), 1, 0, 0, 0, 0, startingDate.Location())
	firstWeekday := int(firstDayOfMonth.Weekday()) // Weekday of the first day (Sunday=0)
	if firstWeekday == 0 {
		firstWeekday = 7 // Adjust Sunday to the last day (Monday-Sunday)
	}

	var cells []fyne.CanvasObject

	// Add empty cells for alignment
	for i := 1; i < firstWeekday; i++ {
		cells = append(cells, canvas.NewRectangle(color.Transparent))
	}

	// Add day cells for the current month
	daysInMonth := daysIn(startingDate)
	for day := 1; day <= daysInMonth; day++ {
		// dayNumber := day // Capturar el número del día para el callback

		// Crear el texto del botón
		text := canvas.NewText(fmt.Sprintf("%d", day), color.Black)
		text.Alignment = fyne.TextAlignCenter

		var bgColor color.Color = color.White
		amountStyle := styles.GetStyleForAmount(0)

		// If not in the future, make it gray
		if day <= today.Day()+1 {
			// Tomar el amount del bloque
			amountForDay := blocks.GetTotalAmountForDay(startingDate, day)
			amountStyle = styles.GetStyleForAmount(amountForDay)

			bgColor = amountStyle.BGColor

		}

		// Determinar el color de fondo
		// bgColor := color.NRGBA{R: 200, G: 255, B: 200, A: 255} // Color por defecto
		if today.Day() == day && today.Month() == startingDate.Month() && today.Year() == startingDate.Year() {
			bgColor = amountStyle.BGColor
			text.TextStyle.Bold = true
		}

		// Crear el fondo como un rectángulo
		bg := canvas.NewRectangle(bgColor)

		// Crear un contenedor para el fondo y el texto
		cell := container.NewStack(
			bg,   // Fondo
			text, // Texto encima
		)

		// // Añadir interacción con clics
		// clickableCell := widget.NewButton("sangra", func() {
		// 	fmt.Printf("Day %d selected\n", dayNumber)
		// })

		// Agregar la celda interactiva a la lista
		cells = append(cells, cell)
	}

	return cells
}

// daysIn calculates the number of days in a month
func daysIn(t time.Time) int {
	nextMonth := t.AddDate(0, 1, 0)
	return int(nextMonth.Sub(t).Hours() / 24)
}

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label

	onSelected func(date time.Time)
}

func MakeCalendar(startingDate time.Time) *xwidget.Calendar {
	i := widget.NewLabel("Please Choose a Date")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel("")
	l.Alignment = fyne.TextAlignCenter
	d := &date{instruction: i, dateChosen: l}

	// Defines which date you would like the calendar to start
	var calendar *xwidget.Calendar
	calendar = xwidget.NewCalendar(startingDate, d.onSelected)

	return calendar
}
