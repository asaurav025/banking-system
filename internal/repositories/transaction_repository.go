package repositories

import (
	"banking-system/internal/models"
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

// For Singleton
var transactionRepositoryOnce sync.Once
var transactionRepositoryInstance *transactionRepository

func NewTransactionRepository(dbInstance *gorm.DB) *transactionRepository {
	if transactionRepositoryInstance == nil {
		transactionRepositoryOnce.Do(func() {
			transactionRepositoryInstance = &transactionRepository{
				db: dbInstance,
			}
		})
	}
	return transactionRepositoryInstance
}

func (repo *transactionRepository) Create(ctx context.Context, item *models.Transaction) (*models.Transaction, error) {
	response := repo.db.Create(&item)
	if response.Error != nil {
		return nil, response.Error
	}
	return item, nil
}

func (repo *transactionRepository) GetTransaction(ctx context.Context, accountId uuid.UUID, from time.Time, to time.Time) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	response := repo.db.Where("(source_id = ? OR destination_id = ?) AND created_on BETWEEN ? AND ?", accountId, accountId, to, from).Find(&transactions)
	if response.Error != nil {
		return nil, response.Error
	}
	return &transactions, nil
}

func (repo *transactionRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string, comment string) error {
	response := repo.db.Model(&models.Transaction{}).Where(models.Transaction{
		Common: models.Common{
			Id: id,
		},
	}).Update(models.Transaction{
		Status:  status,
		Comment: comment,
	})
	if response.Error != nil {
		return response.Error
	}
	return nil
}
