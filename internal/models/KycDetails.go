package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type KycDetails struct {
	Common
	Status       string    `gorm:"column:status"`
	GovtIdNumber string    `gorm:"column:govt_id_number"`
	ExpiryDate   time.Time `gorm:"column:expiry_date"`
	VerifiedBy   string    `gorm:"column:verified_by"`
}

func (model *KycDetails) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", time.Now())
	scope.SetColumn("UpdatedOn", time.Now())
	scope.SetColumn("UpdatedBy", model.CreatedBy)
	return nil
}

func (KycDetails) TableName() string {
	return "kyc_details"
}
