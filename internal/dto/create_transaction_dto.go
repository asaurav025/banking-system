package dto

type CreateTransactionDTO struct {
	Amount      uint   `json:"amount"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}
