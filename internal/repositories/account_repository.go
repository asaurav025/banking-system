package repositories

import (
	"banking-system/internal/models"
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

// For Singleton
var accountRepositoryOnce sync.Once
var accountRepositoryInstance *accountRepository

func NewAccountRepository(dbInstance *gorm.DB) *accountRepository {
	if accountRepositoryInstance == nil {
		accountRepositoryOnce.Do(func() {
			accountRepositoryInstance = &accountRepository{
				db: dbInstance,
			}
		})
	}
	return accountRepositoryInstance
}

func (repo *accountRepository) Create(ctx context.Context, item *models.Account) (*models.Account, error) {
	response := repo.db.Create(&item)
	if response.Error != nil {
		return nil, response.Error
	}
	return item, nil
}

func (repo *accountRepository) Find(ctx context.Context, id uuid.UUID) (*[]models.Account, error) {
	var accounts []models.Account
	acc := new(models.Account)
	acc.Id = id
	response := repo.db.Where(acc).Find(&accounts)
	if response.Error != nil {
		return nil, response.Error
	}
	return &accounts, nil
}

func (repo *accountRepository) UpdateBalance(ctx context.Context, id uuid.UUID, balance uint) error {
	user := ctx.Value("user.id").(string)
	response := repo.db.Model(&models.Account{}).Where(models.Account{
		Common: models.Common{
			Id:        id,
			UpdatedBy: user,
			UpdatedOn: time.Now(),
		},
	}).Update(models.Account{
		Balance: balance,
	})
	if response.Error != nil {
		return response.Error
	}
	return nil
}
