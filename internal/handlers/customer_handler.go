package handlers

import (
	"context"
	"net/http"
	"sync"

	"banking-system/internal/dto"
	"banking-system/internal/interfaces/iservices"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	ICustomerService   iservices.ICustomerService
	IKycDetailsService iservices.IKycDetailsService
}

var CustomerHandlerInstance *CustomerHandler
var CustomerHandlerOnce sync.Once

const USER_ID = "user.id"

func NewCustomerHandler(
	ICustomerService iservices.ICustomerService,
	IKycDetailsService iservices.IKycDetailsService,
) *CustomerHandler {
	if CustomerHandlerInstance == nil {
		CustomerHandlerOnce.Do(func() {
			CustomerHandlerInstance = &CustomerHandler{
				ICustomerService:   ICustomerService,
				IKycDetailsService: IKycDetailsService,
			}
		})
	}
	return CustomerHandlerInstance
}

func (handler *CustomerHandler) AddCustomer(c echo.Context) error {
	log.Info("Method: Add customer")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	requestBody := new(dto.CreateCustomerDTO)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	_, err := handler.ICustomerService.Create(ctx, requestBody)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusCreated)
}

func (handler *CustomerHandler) LinkAccount(c echo.Context) error {
	log.Info("Method: Link accont")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	customerID, err := uuid.Parse(c.Param("customerId"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	accountId, err := uuid.Parse(c.Param("accountId"))
	if err != nil {
		log.Error("Failed to parse account ID. Error: ", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err = handler.ICustomerService.AddAccount(ctx, customerID, accountId)
	if err != nil {
		log.Error("Failed to Link customer. Error: ", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)

}

func (handler *CustomerHandler) UpdateKyc(c echo.Context) error {
	log.Info("Method: Update KYC")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	customerID, err := uuid.Parse(c.Param("customerId"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	requestBody := new(dto.CreateKycDetailsDTO)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	kycDetails, err := handler.IKycDetailsService.Create(ctx, requestBody)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = handler.ICustomerService.UpdateKyc(ctx, customerID, kycDetails.Id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}

func (handler *CustomerHandler) GetCustomer(c echo.Context) error {
	log.Info("Method: Get Customer data")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	custID := c.Param("id")
	customerID, err := uuid.Parse(custID)
	if err != nil {
		log.Error("Failed to parse customerId: ", custID)
		return c.NoContent(http.StatusBadRequest)
	}
	data, err := handler.ICustomerService.GetCustomer(ctx, customerID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, data)
}

func (handler *CustomerHandler) DeleteCustomer(c echo.Context) error {
	log.Info("Method: Delete customer")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))

	customerID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = handler.ICustomerService.Deactivate(ctx, customerID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}
