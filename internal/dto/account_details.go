package dto

import "github.com/google/uuid"

type AccountDetails struct {
	Id      uuid.UUID `json:"account_id"`
	Type    string    `json:"type"`
	Balance string    `json:"balance"`
}
