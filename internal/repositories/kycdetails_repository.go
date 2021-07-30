package repositories

import (
	"banking-system/internal/models"
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type kycDetailsRepository struct {
	db *gorm.DB
}

// For Singleton
var kycDetailsRepositoryOnce sync.Once
var kycDetailsRepositoryInstance *kycDetailsRepository

func NewKycDetailsRepository(dbInstance *gorm.DB) *kycDetailsRepository {
	if kycDetailsRepositoryInstance == nil {
		kycDetailsRepositoryOnce.Do(func() {
			kycDetailsRepositoryInstance = &kycDetailsRepository{
				db: dbInstance,
			}
		})
	}
	return kycDetailsRepositoryInstance
}

func (repo *kycDetailsRepository) Create(ctx context.Context, item *models.KycDetails) (*models.KycDetails, error) {
	response := repo.db.Create(&item)
	if response.Error != nil {
		return nil, response.Error
	}
	return item, nil
}

func (repo *kycDetailsRepository) Get(ctx context.Context, id uuid.UUID) (*[]models.KycDetails, error) {
	var kycDetails []models.KycDetails
	cust := new(models.KycDetails)
	cust.Id = id
	response := repo.db.Where(cust).Find(&kycDetails)
	if response.Error != nil {
		return nil, response.Error
	}
	return &kycDetails, nil
}
func (repo *kycDetailsRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return nil
}
func (repo *kycDetailsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (repo *kycDetailsRepository) Verify(ctx context.Context, id uuid.UUID, verifier string) error {
	return nil
}
