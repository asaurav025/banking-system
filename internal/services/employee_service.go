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
	userEmail := ctx.Value("user.id").(string)
	emps, err := service.iEmployeeRepository.FindByEmail(ctx, userEmail)
	if err != nil {
		log.Error("Failed to fetch employee")
		return nil, err
	}
	if len(*emps) == 0 {
		log.Error("No employee found")
		return nil, errors.New("no employee found")
	}
	emp0 := (*emps)[0]

	if emp0.Type != "ADMIN" {
		log.Error("employee not authorized")
		return nil, errors.New("not authorized")
	}

	employee := new(models.Employee)
	employee.Id = uuid.New()
	employee.Type = "GENERAL"
	employee.CreatedBy = userEmail
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
	userEmail := ctx.Value("user.id").(string)
	emps, err := service.iEmployeeRepository.FindByEmail(ctx, userEmail)
	if err != nil {
		log.Error("Failed to fetch employee")
		return err
	}
	if len(*emps) == 0 {
		log.Error("No employee found")
		return errors.New("no employee found")
	}
	emp0 := (*emps)[0]

	if emp0.Type != "ADMIN" {
		return errors.New("not authorized")
	}

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
		return "", errors.New("employee not found")
	}
	emp := (*emps)[0]

	if b64.StdEncoding.EncodeToString([]byte(password)) != emp.Password {
		log.Error("Wrong credentilas")
		return "", errors.New("wrong credentials")
	}
	return emp.Id.String(), nil
}
