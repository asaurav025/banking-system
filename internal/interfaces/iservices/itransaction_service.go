package iservices

import (
	"banking-system/internal/dto"
	"context"
	"time"

	"github.com/google/uuid"
)

type ITransactionService interface {
	Create(ctx context.Context, item *dto.CreateTransactionDTO) (interface{}, error)
	UpdateStatus(ctx context.Context, id uuid.UUID) error
	GetTransaction(ctx context.Context, accountId uuid.UUID, from time.Time, to time.Time) (string, error)
	AddInterest(ctx context.Context, id uuid.UUID) error
}
