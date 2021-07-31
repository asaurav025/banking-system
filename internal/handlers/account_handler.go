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

type AccountHandler struct {
	IAccountService iservices.IAccountService
}

var AccountHandlerInstance *AccountHandler
var AccountHandlerOnce sync.Once

func NewAccountHandler(
	IAccountService iservices.IAccountService,
) *AccountHandler {
	if AccountHandlerInstance == nil {
		AccountHandlerOnce.Do(func() {
			AccountHandlerInstance = &AccountHandler{
				IAccountService: IAccountService,
			}
		})
	}
	return AccountHandlerInstance
}

func (handler *AccountHandler) CreateAccount(c echo.Context) error {
	log.Info("Method: Create account")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	requestBody := new(dto.AccountRequestDto)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	acc, err := handler.IAccountService.Create(ctx, requestBody)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, acc)
}

func (handler *AccountHandler) GetBalance(c echo.Context) error {
	log.Info("Method: Get balance")
	ctx := context.WithValue(httpContext, USER_ID, c.Get(USER_ID))
	accountId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	acc, err := handler.IAccountService.GetAccount(ctx, accountId)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, &struct {
		Balance uint `json:"balance"`
	}{
		Balance: acc.Balance,
	})
}
