package dto

import "time"

type TransactionReport struct {
	AccountID    string
	From         time.Time
	To           time.Time
	Transactions []Transaction
}

type Transaction struct {
	TransactionID string
	Type          string
	AccountID     string
	Amount        string
	Time          time.Time
}
