package iservices

import (
	"banking-system/internal/dto"
	"context"

	"github.com/google/uuid"
)

type IEmployeeService interface {
	AddEmployee(ctx context.Context, item *dto.EmployeeRequestDto) (interface{}, error)
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
}
