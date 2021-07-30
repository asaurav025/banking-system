package irepositories

import (
	"banking-system/internal/models"
	"context"

	"github.com/google/uuid"
)

type IKycDetailsRepository interface {
	Create(ctx context.Context, item *models.KycDetails) (*models.KycDetails, error)
	Get(ctx context.Context, id uuid.UUID) (*[]models.KycDetails, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	Delete(ctx context.Context, id uuid.UUID) error
	Verify(ctx context.Context, id uuid.UUID, verifier string) error
}
