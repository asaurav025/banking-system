package irepositories

import (
	"banking-system/internal/models"
	"context"

	"github.com/google/uuid"
)

type IEmployeeRepository interface {
	Create(ctx context.Context, item *models.Employee) (*models.Employee, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
