package iservices

import (
	"banking-system/internal/dto"
	"context"

	"github.com/google/uuid"
)

type ICustomerService interface {
	Create(ctx context.Context, item *dto.CreateCustomerDTO) (interface{}, error)
	GetCustomer(ctx context.Context, id uuid.UUID) (dto.GetCustomerDetailsDTO, error)
	UpdateName(ctx context.Context, name string) error
	Deactivate(ctx context.Context, id uuid.UUID) error
	AddAccount(ctx context.Context, id uuid.UUID, accountId uuid.UUID) error
	RemoveAccount(ctx context.Context, id uuid.UUID, accountId uuid.UUID) error
	UpdateKyc(ctx context.Context, id uuid.UUID, kycId uuid.UUID) error
}
