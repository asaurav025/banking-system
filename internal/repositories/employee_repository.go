package repositories

import (
	"banking-system/internal/models"
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

// For Singleton
var employeeRepositoryOnce sync.Once
var employeeRepositoryInstance *employeeRepository

func NewEmployeeRepository(dbInstance *gorm.DB) *employeeRepository {
	if employeeRepositoryInstance == nil {
		employeeRepositoryOnce.Do(func() {
			employeeRepositoryInstance = &employeeRepository{
				db: dbInstance,
			}
		})
	}
	return employeeRepositoryInstance
}

func (repo *employeeRepository) Create(ctx context.Context, item *models.Employee) (*models.Employee, error) {
	response := repo.db.Create(&item)
	if response.Error != nil {
		return nil, response.Error
	}
	return item, nil
}

func (repo *employeeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	emp := new(models.Employee)
	emp.Id = id
	response := repo.db.Delete(emp)
	return response.Error
}
