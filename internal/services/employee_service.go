package services

import (
	"context"
	b64 "encoding/base64"
	"errors"
	"sync"

	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	employee.Password = b64.StdEncoding.EncodeToString([]byte(item.Password))
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

func (service *employeeService) VerifyEmployee(ctx context.Context, email string, password string) (string, error) {
	emps, err := service.iEmployeeRepository.FindByEmail(ctx, email)
	if err != nil {
		log.Error("Failed to get employee")
		return "", err
	}
	if len(*emps) == 0 {
		log.Error("Failed to get employee")
		return "", errors.New("Employee not found")
	}
	emp := (*emps)[0]

	if b64.StdEncoding.EncodeToString([]byte(password)) != emp.Password {
		log.Error("Wrong credentilas")
		return "", errors.New("Wrong credentials")
	}
	return emp.Id.String(), nil
}
