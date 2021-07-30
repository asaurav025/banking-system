package irepositories

import (
	"banking-system/internal/models"
	"context"

	"github.com/google/uuid"
)

type IAccountRepository interface {
	Create(ctx context.Context, item *models.Account) (*models.Account, error)
	Find(ctx context.Context, id uuid.UUID) (*[]models.Account, error)
	UpdateBalance(ctx context.Context, id uuid.UUID, balance uint) error
}
