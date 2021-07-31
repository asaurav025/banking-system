package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	Common
	Type          string    `gorm:"column:type"`
	Amount        uint      `gorm:"column:amount"`
	Unit          string    `gorm:"column:unit"`
	Status        string    `gorm:"column:status"`
	SourceID      uuid.UUID `gorm:"column:source_id"`
	DestinationId uuid.UUID `gorm:"column:destination_id"`
	Comment       string    `gorm:"column:comment"`
}

func (model *Transaction) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", time.Now())
	scope.SetColumn("UpdatedOn", time.Now())
	scope.SetColumn("UpdatedBy", model.CreatedBy)
	return nil
}

func (Transaction) TableName() string {
	return "transaction"
}
