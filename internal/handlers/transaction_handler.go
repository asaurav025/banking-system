package handlers

import (
	"banking-system/internal/dto"
	"banking-system/internal/interfaces/iservices"
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	ITransactionService iservices.ITransactionService
}

var TransactionHandlerInstance *TransactionHandler
var TransactionHandlerOnce sync.Once

func NewTransactionHandler(
	ITransactionService iservices.ITransactionService,
) *TransactionHandler {
	if TransactionHandlerInstance == nil {
		TransactionHandlerOnce.Do(func() {
			TransactionHandlerInstance = &TransactionHandler{
				ITransactionService: ITransactionService,
			}
		})
	}
	return TransactionHandlerInstance
}

func (handler *TransactionHandler) CreateTransaction(c echo.Context) error {
	log.Info("Method: Create transaction")
	ctx := context.WithValue(httpContext, USER_ID, "")
	requestBody := new(dto.CreateTransactionDTO)
	bindErr := c.Bind(requestBody)
	if bindErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	trans, err := handler.ITransactionService.Create(ctx, requestBody)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, trans)
}

func (handler *TransactionHandler) GetTransaction(c echo.Context) error {
	log.Info("Mehtod: Get Transaction")
	ctx := context.WithValue(httpContext, USER_ID, "")
	accountID, err := uuid.Parse(c.Param("accountId"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	layout := "2006-01-02T15:04:05.000Z"
	from, err := time.Parse(layout, c.Param("from"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	to, err := time.Parse(layout, c.Param("to"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	report, err := handler.ITransactionService.GetTransaction(ctx, accountID, from, to)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.File(report)
}
