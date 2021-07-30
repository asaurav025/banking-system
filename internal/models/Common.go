package models

import (
	"time"

	"github.com/google/uuid"
)

type Common struct {
	Id        uuid.UUID `gorm:"primary_key" column:"id"`
	CreatedOn time.Time `gorm:"column:created_on"`
	UpdatedOn time.Time `gorm:"column:updated_on"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedBy string    `gorm:"column:updated_by"`
}
