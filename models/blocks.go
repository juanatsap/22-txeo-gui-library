package models

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
	"time"
	"txeo-gui-library/styles"

	"fyne.io/fyne/v2/widget"
	"github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

type Block struct {
	Concept         Concept
	ConceptAsString string
	Date            string
	Concept2        string
	Amount          float64
	Balance         string
	Category        Category
}
type Blocks []Block

func (b *Block) NewBlock(conceptAsString string, date string, concept2 string, amount float64, balance string) *Block {

	concept := NewConceptFromString(conceptAsString)
	return &Block{Concept: concept, Date: date, Concept2: concept2, Amount: amount, Balance: balance}
}
func (b *Blocks) Len() int           { return len(*b) }
func (b *Blocks) Swap(i, j int)      { (*b)[i], (*b)[j] = (*b)[j], (*b)[i] }
func (b *Blocks) Less(i, j int) bool { return (*b)[i].Date < (*b)[j].Date }

func (b Block) PrintInfo() {
	log.Infof("Object: %#v", b)
}
func (b Block) GetBackgroundGlobalStyle() *widget.CustomTextGridStyle {
	greenStyle := &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}}
	redStyle := &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 192, G: 64, B: 64, A: 128}}

	amountStyle := styles.GetStyleForAmount(b.Amount)

	var returnedStyle *widget.CustomTextGridStyle
	if b.Amount > 0 {
		returnedStyle = amountStyle
	} else {
		returnedStyle = greenStyle
	}

	// If category is nil, return redStyle
	if b.Category.Name == "" {
		return redStyle
	}

	return returnedStyle
}
func (b Block) GetDateStyle(blocks Blocks) *widget.CustomTextGridStyle {

	// take the date of the block, and for all the blocks in that date, sum the amounts
	oneDayAmount := 0.0
	for i := 0; i < len(blocks); i++ {
		if blocks[i].Date == b.Date {
			oneDayAmount += blocks[i].Amount
		}
	}
	oneDateStyle := styles.GetStyleForAmount(oneDayAmount)

	return oneDateStyle
}
func (b Block) GetAmountStyle() *widget.CustomTextGridStyle {
	return styles.GetStyleForAmount(b.Amount)
}
func (b Block) GetBalanceStyle() *widget.CustomTextGridStyle {
	return styles.GetStyleForBalance(b.GetBalanceAsFloat())
}
func (b Block) GetConceptStyle() *widget.CustomTextGridStyle {
	return styles.GetStyleForAmount(0)
}
func (b Block) GetBalanceAsFloat() float64 {
	balance, err := ParseBalanceString(b.Balance)
	if err != nil {
		return 0
	}
	return balance
}
func (b Block) Println() {

	fmt.Print("\n--------------------------------------------------------------------------------------------------------------------\n")
	log.Infof("  📅  %-12s %s %-15s ✏️ [%20s  ] 💵 Amount: %4.2f 💵 Balance: %9s", aurora.BrightWhite(b.Date), b.Category.Icon, aurora.BrightYellow(b.Category.ShortName), aurora.BrightWhite(b.Concept.Name), aurora.BrightRed(b.Amount), aurora.BrightGreen(b.Balance))
	fmt.Print("--------------------------------------------------------------------------------------------------------------------\n")
}
func (b Block) PrintlnForClick(row int, direction string) {

	fmt.Print("\n---------------------------------------------------------------------------------------------------------------------------------------------\n")
	log.Infof("   🮰 %5s-clicked row: %d -->  📅  %-12s %s %-15s ✏️ [%20s  ] 💵 Amount: %4.2f 💵 Balance: %9s", strings.ToTitle(direction), row, aurora.BrightWhite(b.Date), b.Category.Icon, aurora.BrightYellow(b.Category.ShortName), aurora.BrightWhite(b.Concept.Name), aurora.BrightRed(b.Amount), aurora.BrightGreen(b.Balance))
	fmt.Print("---------------------------------------------------------------------------------------------------------------------------------------------\n")
}
func (b Block) GetAmountAsFloat() float64 {
	return b.Amount
}
func (b *Blocks) AddBlock(block Block) {

	*b = append(*b, block)
}
func (b *Block) AssignCategoryForBlock(categories Categories) {

	// Get Unknown Category
	var category Category
	category = category.TryToAssignCategory(*b, categories)

	// Try to assign a category to the block

	// Assign the category to the block
	b.Category = category
}
func (b *Block) GetCategory() Category {
	return b.Category
}

/* ╭──────────────────────────────────────────╮ */
/* │                  BLOCKS                  │ */
/* ╰──────────────────────────────────────────╯ */
func (b Blocks) GetAmountAsFloat() float64 {

	totalAmount := 0.0
	for i := 0; i < len(b); i++ {
		totalAmount += b[i].Amount
	}
	return totalAmount
}
func ParseBalanceString(balanceStr string) (float64, error) {
	// Quitar espacios al principio y al final
	s := strings.TrimSpace(balanceStr)

	// Eliminar el símbolo de euro si existe al final
	s = strings.TrimSuffix(s, "€")
	s = strings.TrimSpace(s) // Quitar espacios que pudieran quedar

	// Revisar el signo
	sign := 1.0
	if strings.HasPrefix(s, "-") {
		sign = -1.0
		s = s[1:] // Quitar el signo
	} else if strings.HasPrefix(s, "+") {
		s = s[1:] // Quitar el signo '+'
	}

	s = strings.TrimSpace(s)

	// Reemplazar coma por punto (asumiendo formato español)
	s = strings.ReplaceAll(s, ",", ".")

	// Ahora s debería ser un número válido en formato inglés, ej. "100.96" o "20.15"
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return sign * value, nil
}

func (b Blocks) GetTotalAmountForDay(startingDate time.Time, day int) float64 {

	totalAmount := 0.0
	for i := 0; i < len(b); i++ {
		if b[i].Date == startingDate.Format("2006-01-02") {
			totalAmount += b[i].Amount
		}
	}
	return totalAmount
}
