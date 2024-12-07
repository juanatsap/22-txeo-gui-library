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

	// Rango para ingresos positivos
	minIncomeVal     = 0.0
	maxIncomeVal     = 3000.0
	startGreenColor  = color.NRGBA{R: 0xD4, G: 0xF8, B: 0xD4, A: 0xFF} // #D4F8D4 (Verde pastel)
	endGreenColor    = color.NRGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF} // #00FF00 (Verde intenso)
	negativeRedColor = color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF} // #FF0000 (Rojo para negativos)

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

func GetStyleForBalance(balance float64) *widget.CustomTextGridStyle {
	var bgColor color.NRGBA

	if balance < 0 {
		// Caso negativo: fondo rojo fijo
		bgColor = negativeRedColor
	} else {
		// Caso positivo o cero: interpolar entre startGreenColor y endGreenColor
		val := balance
		if val > maxIncomeVal {
			val = maxIncomeVal // Limitar por arriba
		}

		fraction := (val - minIncomeVal) / (maxIncomeVal - minIncomeVal)
		if fraction < 0 {
			fraction = 0
		} else if fraction > 1 {
			fraction = 1
		}

		R := float64(startGreenColor.R) + fraction*(float64(endGreenColor.R)-float64(startGreenColor.R))
		G := float64(startGreenColor.G) + fraction*(float64(endGreenColor.G)-float64(startGreenColor.G))
		B := float64(startGreenColor.B) + fraction*(float64(endGreenColor.B)-float64(startGreenColor.B))
		A := float64(startGreenColor.A) + fraction*(float64(endGreenColor.A)-float64(startGreenColor.A))

		bgColor = color.NRGBA{
			R: uint8(math.Round(R)),
			G: uint8(math.Round(G)),
			B: uint8(math.Round(B)),
			A: uint8(math.Round(A)),
		}
	}

	// Calcular luminancia para determinar el color de texto
	luminance := (0.299*float64(bgColor.R) + 0.587*float64(bgColor.G) + 0.114*float64(bgColor.B)) / 255.0

	var fgColor color.NRGBA
	if luminance > 0.5 {
		// Fondo claro, texto oscuro
		fgColor = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
	} else {
		// Fondo oscuro, texto claro
		fgColor = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	}

	return &widget.CustomTextGridStyle{FGColor: &fgColor, BGColor: &bgColor}
}
