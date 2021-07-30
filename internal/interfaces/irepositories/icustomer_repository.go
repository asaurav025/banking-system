package irepositories

import (
	"context"

	"banking-system/internal/models"

	"github.com/google/uuid"
)

type ICustomerRepository interface {
	Create(ctx context.Context, item *models.Customer) (*models.Customer, error)
	Find(ctx context.Context, id uuid.UUID) (*[]models.Customer, error)
	Update(ctx context.Context, item *models.Customer) (*models.Customer, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
