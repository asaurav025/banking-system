package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	Common
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Type  string `gorm:"column:type"`
}

func (model *Employee) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", time.Now())
	scope.SetColumn("UpdatedOn", time.Now())
	scope.SetColumn("UpdatedBy", model.CreatedBy)
	return nil
}

func (Employee) TableName() string {
	return "employee"
}
