package services

import (
	"context"
	"sync"
	"time"

	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"

	"github.com/google/uuid"
)

type kycDetialsService struct {
	IKycDetailsRepository irepositories.IKycDetailsRepository
}

// For Singleton
var kycDetialsServiceOnce sync.Once
var kycDetialsServiceInstance *kycDetialsService

func NewKycDetialsService(
	kycDetailsRepo irepositories.IKycDetailsRepository,
) *kycDetialsService {
	if kycDetialsServiceInstance == nil {
		kycDetialsServiceOnce.Do(func() {
			kycDetialsServiceInstance = &kycDetialsService{
				IKycDetailsRepository: kycDetailsRepo,
			}
		})
	}
	return kycDetialsServiceInstance
}

func (service *kycDetialsService) Create(ctx context.Context, item *dto.CreateKycDetailsDTO) (*models.KycDetails, error) {
	kyc := new(models.KycDetails)
	kyc.Id = uuid.New()
	kyc.GovtIdNumber = item.GovtIDNumber
	kyc.CreatedBy = ctx.Value("user.id").(string)
	kyc.Status = "CREATED"
	layout := "2006-01-02"
	t, err := time.Parse(layout, item.ExpiryDate)
	if err != nil {
		return nil, err
	}
	kyc.ExpiryDate = t
	_, err = service.IKycDetailsRepository.Create(ctx, kyc)
	if err != nil {
		return nil, err
	}
	return kyc, nil
}
func (service *kycDetialsService) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return nil
}
func (service *kycDetialsService) Delete(ctx context.Context, id uuid.UUID) error {
	return service.IKycDetailsRepository.Delete(ctx, id)
}
