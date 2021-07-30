package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	Common
	Type    string `gorm:"column:type"`
	Balance uint   `gorm:"column:balance"`
	Unit    string `gorm:"column:unit"`
}

func (model *Account) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", time.Now())
	scope.SetColumn("UpdatedOn", time.Now())
	scope.SetColumn("UpdatedBy", model.CreatedBy)
	return nil
}
