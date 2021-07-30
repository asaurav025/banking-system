package dto

type CreateKycDetailsDTO struct {
	GovtIDNumber string `json:"govt_id_number"`
	ExpiryDate   string `json:"expiry_date"`
}
