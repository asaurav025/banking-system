package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	Common
	Name                      string          `gorm:"column:name"`
	KycDetailsId              uuid.UUID       `gorm:"column:kyc_detials_id"`
	AccountDetails            string          `gorm:"column:account_detials" sql:"type:JSONB NOT BULL DEFAULT '{}'::JSONB"`
	AccountDetailsMap         jsonMap         `gorm:"-"`
	AccountDetailsKeyValueMap jsonKeyValueMap `gorm:"-"`
}

type jsonMap map[string]string
type jsonKeyValueMap map[string]interface{}

func (model *Customer) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", time.Now())
	scope.SetColumn("UpdatedOn", time.Now())
	scope.SetColumn("UpdatedBy", model.CreatedBy)
	return nil
}

func (model *Customer) AfterFind() error {
	err := json.Unmarshal([]byte(model.AccountDetails), &model.AccountDetailsMap)
	if err != nil {
		err = json.Unmarshal([]byte(model.AccountDetails), &model.AccountDetailsKeyValueMap)
	}
	return err
}

func (Customer) TableName() string {
	return "customer"
}
