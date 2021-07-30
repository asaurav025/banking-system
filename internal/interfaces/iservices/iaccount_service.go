package iservices

import (
	"banking-system/internal/dto"
	"banking-system/internal/models"
	"context"

	"github.com/google/uuid"
)

type IAccountService interface {
	Create(ctx context.Context, item *dto.AccountRequestDto) (*models.Account, error)
	GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error)
	UpdateBalance(ctx context.Context, id uuid.UUID, balance uint) error
}
