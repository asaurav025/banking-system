package iservices

import (
	"banking-system/internal/dto"
	"banking-system/internal/models"
	"context"

	"github.com/google/uuid"
)

type IKycDetailsService interface {
	Create(ctx context.Context, item *dto.CreateKycDetailsDTO) (*models.KycDetails, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	Delete(ctx context.Context, id uuid.UUID) error
}
