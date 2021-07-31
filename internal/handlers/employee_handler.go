package handlers

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"banking-system/internal/dto"
	"banking-system/internal/interfaces/iservices"
	"banking-system/pkg/jwt"

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
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	requestBody := new(dto.EmployeeRequestDto)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	_, err := handler.IEmployeeService.AddEmployee(ctx, requestBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (handler *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	log.Info("Method: Delete employee")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))

	employeeID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = handler.IEmployeeService.DeleteEmployee(ctx, employeeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (handler *EmployeeHandler) Login(c echo.Context) error {
	requestBody := new(dto.LoginDTO)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		log.Error("Failed to parse body")
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}
	_, err := handler.IEmployeeService.VerifyEmployee(httpContext, requestBody.Email, requestBody.Password)
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}
	var response dto.JwtCreationResponse
	token, err := jwt.CreateToken(requestBody.Email)
	if err != nil {
		log.Error("Failed to generate token")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	response.Duration = fmt.Sprint(jwt.EXPIRES_AFTER)
	response.Token = token

	return c.JSON(http.StatusCreated, response)
}
