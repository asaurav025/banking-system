package irepositories

import (
	"banking-system/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
)

type ITrasactionRepository interface {
	Create(ctx context.Context, item *models.Transaction) (*models.Transaction, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string, comment string) error
	GetTransaction(ctx context.Context, accountId uuid.UUID, from time.Time, to time.Time) (*[]models.Transaction, error)
}
