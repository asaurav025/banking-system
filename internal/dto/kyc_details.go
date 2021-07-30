package dto

import "github.com/google/uuid"

type KycDetails struct {
	Id           uuid.UUID `json:"kyc_id"`
	Status       string    `json:"status"`
	GovtIDNumber string    `json:"govt_id_number"`
	ExpiryDate   string    `json:"expiry_date"`
	VerifiedBy   string    `json:"verified_by"`
}
