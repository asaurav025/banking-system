package repositories

import (
	"banking-system/internal/models"
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

// For Singleton
var customerRepositoryOnce sync.Once
var customerRepositoryInstance *customerRepository

func NewCustomerRepository(dbInstance *gorm.DB) *customerRepository {
	if customerRepositoryInstance == nil {
		customerRepositoryOnce.Do(func() {
			customerRepositoryInstance = &customerRepository{
				db: dbInstance,
			}
		})
	}
	return customerRepositoryInstance
}

func (repo *customerRepository) Create(ctx context.Context, item *models.Customer) (*models.Customer, error) {
	response := repo.db.Create(&item)
	if response.Error != nil {
		return nil, response.Error
	}
	return item, nil
}
func (repo *customerRepository) Find(ctx context.Context, id uuid.UUID) (*[]models.Customer, error) {
	var customers []models.Customer
	cust := new(models.Customer)
	cust.Id = id
	response := repo.db.Where(cust).Find(&customers)
	if response.Error != nil {
		return nil, response.Error
	}
	return &customers, nil
}
func (repo *customerRepository) Update(ctx context.Context, item *models.Customer) (*models.Customer, error) {
	response := repo.db.Model(&models.Customer{}).Where(models.Customer{
		Common: models.Common{
			Id:        item.Id,
			UpdatedOn: time.Now(),
		},
	}).Update(
		item,
	)
	if response.Error != nil {
		return nil, response.Error
	}
	return item, nil
}
func (repo *customerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	cust := new(models.Customer)
	cust.Id = id
	response := repo.db.Delete(cust)
	return response.Error
}
