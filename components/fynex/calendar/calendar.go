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

	// Use the selectedDate parameter instead of time.Now()
	startingDate := selectedDate
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

// MakeCustomCalendarWithCallback creates a fully custom calendar with day selection callback
func MakeCustomCalendarWithCallback(selectedDate time.Time, blocks models.Blocks, onDaySelected func(int)) *fyne.Container {

	// Use the selectedDate parameter instead of time.Now()
	startingDate := selectedDate
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

	// Generate calendar cells with callback
	cells := createCustomCalendarCellsWithCallback(startingDate, blocks, onDaySelected)

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

// MakeCustomCalendarWithCallbackAndSelection creates a calendar with day selection callback and highlights selected day
func MakeCustomCalendarWithCallbackAndSelection(selectedDate time.Time, blocks models.Blocks, selectedDay int, onDaySelected func(int)) *fyne.Container {

	// Use the selectedDate parameter instead of time.Now()
	startingDate := selectedDate
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

	// Generate calendar cells with callback and selection
	cells := createCustomCalendarCellsWithCallbackAndSelection(startingDate, blocks, selectedDay, onDaySelected)

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

// createCustomCalendarCellsWithCallback generates the cells with click handling
func createCustomCalendarCellsWithCallback(startingDate time.Time, blocks models.Blocks, onDaySelected func(int)) []fyne.CanvasObject {
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
		dayNumber := day // Capture the day number for the callback

		var bgColor color.Color = color.White
		amountStyle := styles.GetStyleForAmount(0)

		// If not in the future, make it gray
		if day <= today.Day()+1 {
			// Get the amount for the day
			amountForDay := blocks.GetTotalAmountForDay(startingDate, day)
			amountStyle = styles.GetStyleForAmount(amountForDay)
			bgColor = amountStyle.BGColor
		}

		// Check if this is today
		isToday := today.Day() == day && today.Month() == startingDate.Month() && today.Year() == startingDate.Year()

		// Determine background color for today
		if isToday {
			bgColor = amountStyle.BGColor
		}

		// Create the background rectangle
		bg := canvas.NewRectangle(bgColor)

		// Make the cell clickable using a button with the day number
		// Capture the day number properly to avoid closure issues
		capturedDay := dayNumber
		clickableCell := &widget.Button{
			Text: fmt.Sprintf("%d", dayNumber),
			OnTapped: func() {
				if onDaySelected != nil {
					onDaySelected(capturedDay)
				}
			},
			Importance: widget.LowImportance, // Make it blend with the background
		}

		// If it's today, make it bold (we'll need to use a different approach since Button doesn't support text style directly)
		if isToday {
			clickableCell.Importance = widget.HighImportance // This will make it stand out more
		}

		// Style the button to look like the original calendar cell
		// We'll use the background color we calculated
		styledButton := container.NewStack(
			bg, // Use the same background
			container.NewCenter(clickableCell),
		)

		// Add the interactive cell to the list
		cells = append(cells, styledButton)
	}

	return cells
}

// createCustomCalendarCellsWithCallbackAndSelection generates cells with click handling and selection highlighting
func createCustomCalendarCellsWithCallbackAndSelection(startingDate time.Time, blocks models.Blocks, selectedDay int, onDaySelected func(int)) []fyne.CanvasObject {
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
		dayNumber := day // Capture the day number for the callback

		var bgColor color.Color = color.White
		amountStyle := styles.GetStyleForAmount(0)

		// If not in the future, make it gray
		if day <= today.Day()+1 {
			// Get the amount for the day
			amountForDay := blocks.GetTotalAmountForDay(startingDate, day)
			amountStyle = styles.GetStyleForAmount(amountForDay)
			bgColor = amountStyle.BGColor
		}

		// Check if this is today
		isToday := today.Day() == day && today.Month() == startingDate.Month() && today.Year() == startingDate.Year()

		// Check if this is the selected day
		isSelected := day == selectedDay

		// Determine background color
		if isToday {
			bgColor = amountStyle.BGColor
		}

		// If this is the selected day, add a border or modify the background
		if isSelected {
			// Make the selected day stand out with a darker/highlighted background
			if bgColor != nil {
				// Darken the existing color slightly for selection
				if nrgba, ok := bgColor.(*color.NRGBA); ok {
					bgColor = &color.NRGBA{
						R: uint8(float64(nrgba.R) * 0.8),
						G: uint8(float64(nrgba.G) * 0.8),
						B: uint8(float64(nrgba.B) * 0.8),
						A: nrgba.A,
					}
				}
			} else {
				// Use a light blue for selection if no other color
				bgColor = &color.NRGBA{R: 100, G: 150, B: 255, A: 255}
			}
		}

		// Create the background rectangle
		bg := canvas.NewRectangle(bgColor)

		// Make the cell clickable using a button with the day number
		// Capture the day number properly to avoid closure issues
		capturedDay := dayNumber
		clickableCell := &widget.Button{
			Text: fmt.Sprintf("%d", dayNumber),
			OnTapped: func() {
				if onDaySelected != nil {
					onDaySelected(capturedDay)
				}
			},
			Importance: widget.LowImportance, // Make it blend with the background
		}

		// If it's today or selected, make it stand out more
		if isToday || isSelected {
			clickableCell.Importance = widget.HighImportance
		}

		// Style the button to look like the original calendar cell
		styledButton := container.NewStack(
			bg, // Use the same background
			container.NewCenter(clickableCell),
		)

		// Add the interactive cell to the list
		cells = append(cells, styledButton)
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
	d := &date{}

	// Defines which date you would like the calendar to start
	var calendar *xwidget.Calendar
	calendar = xwidget.NewCalendar(startingDate, d.onSelected)

	return calendar
}
