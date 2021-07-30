package handlers

import (
	"banking-system/internal/dto"
	"banking-system/internal/interfaces/iservices"
	"context"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type EmployeeHandler struct {
	IEmployeeService iservices.IEmployeeService
}

var EmployeeHandlerInstance *EmployeeHandler
var EmployeeHandlerOnce sync.Once

func NewEmployeeHandler(
	IEmployeeService iservices.IEmployeeService,
) *EmployeeHandler {
	if EmployeeHandlerInstance == nil {
		EmployeeHandlerOnce.Do(func() {
			EmployeeHandlerInstance = &EmployeeHandler{
				IEmployeeService: IEmployeeService,
			}
		})
	}
	return EmployeeHandlerInstance
}

var httpContext = context.Background()

func (handler *EmployeeHandler) AddEmployee(c echo.Context) error {
	log.Info("Method: Add employee")
	ctx := context.WithValue(httpContext, "user.id", "")
	requestBody := new(dto.EmployeeRequestDto)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	_, err := handler.IEmployeeService.AddEmployee(ctx, requestBody)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusCreated)
}

func (handler *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	log.Info("Method: Delete employee")
	ctx := context.WithValue(httpContext, "user.id", "")

	employeeID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = handler.IEmployeeService.DeleteEmployee(ctx, employeeID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}
