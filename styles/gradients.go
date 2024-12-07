package styles

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2/widget"
)

// Definimos los valores mínimos y máximos y los colores inicial y final
var (
	minVal     = 5.0
	maxVal     = 1000.0
	startColor = color.NRGBA{R: 0xF8, G: 0xD4, B: 0x95, A: 0xFF} // #F8D495
	endColor   = color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF} // #FF0000
)

func GetStyleForAmount(amount float64) *widget.CustomTextGridStyle {
	// Ajustar el amount al rango [minVal, maxVal]
	if amount < minVal {
		amount = minVal
	}
	if amount > maxVal {
		amount = maxVal
	}

	// Redondear a múltiplos de 5
	snapVal := 5.0 * math.Round(amount/5.0)

	// Calcular la fracción en el rango
	fraction := (snapVal - minVal) / (maxVal - minVal)
	if fraction < 0 {
		fraction = 0
	} else if fraction > 1 {
		fraction = 1
	}

	// Interpolar cada componente
	R := float64(startColor.R) + fraction*(float64(endColor.R)-float64(startColor.R))
	G := float64(startColor.G) + fraction*(float64(endColor.G)-float64(startColor.G))
	B := float64(startColor.B) + fraction*(float64(endColor.B)-float64(startColor.B))
	A := float64(startColor.A) + fraction*(float64(endColor.A)-float64(startColor.A))

	c := color.NRGBA{
		R: uint8(math.Round(R)),
		G: uint8(math.Round(G)),
		B: uint8(math.Round(B)),
		A: uint8(math.Round(A)),
	}

	// Calcular luminancia para determinar el color del texto
	luminance := (0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B)) / 255.0

	var frontendColor color.NRGBA
	if luminance > 0.5 {
		// Fondo claro, texto oscuro
		frontendColor = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
	} else {
		// Fondo oscuro, texto claro
		frontendColor = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	}

	return &widget.CustomTextGridStyle{FGColor: &frontendColor, BGColor: &c}
}
