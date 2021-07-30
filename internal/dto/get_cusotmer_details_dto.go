package dto

import "github.com/google/uuid"

type GetCustomerDetailsDTO struct {
	Id             uuid.UUID        `json:"customer_id"`
	Name           string           `json:"name"`
	AccountDetails []AccountDetails `json:"account_details"`
	KycDetails     KycDetails       `json:"kyc_details"`
}
