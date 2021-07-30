package services

import (
	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"
	"context"
	"sync"

	"github.com/google/uuid"
)

type employeeService struct {
	iEmployeeRepository irepositories.IEmployeeRepository
}

// For Singleton
var employeeServiceOnce sync.Once
var employeeServiceInstance *employeeService

func NewEmployeeService(
	employeeRepo irepositories.IEmployeeRepository,
) *employeeService {
	if employeeServiceInstance == nil {
		employeeServiceOnce.Do(func() {
			employeeServiceInstance = &employeeService{
				iEmployeeRepository: employeeRepo,
			}
		})
	}
	return employeeServiceInstance
}

func (service *employeeService) AddEmployee(ctx context.Context, item *dto.EmployeeRequestDto) (interface{}, error) {
	// Todo : add check for admin

	employee := new(models.Employee)
	employee.Id = uuid.New()
	employee.Type = "GENERAL"
	employee.CreatedBy = ctx.Value("user.id").(string)
	employee.Email = item.Email
	employee.Name = item.Name
	emp, err := service.iEmployeeRepository.Create(ctx, employee)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (service *employeeService) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	// Todo: add check for admin

	return service.iEmployeeRepository.Delete(ctx, id)
}
