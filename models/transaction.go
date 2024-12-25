package models

type Transaction struct {
	Concept     string
	Amount      string
	Date        string
	Description string
}
type Transactions []Transaction
