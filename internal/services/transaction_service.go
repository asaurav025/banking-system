package services

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type transactionService struct {
	iTransactionRepository irepositories.ITrasactionRepository
	iAccountRepository     irepositories.IAccountRepository
}

// For Singleton
var transactionServiceOnce sync.Once
var transactionServiceInstance *transactionService

const RATE = 0.035

func NewTransactionService(
	transactionRepo irepositories.ITrasactionRepository,
	accountRepo irepositories.IAccountRepository,
) *transactionService {
	if transactionServiceInstance == nil {
		transactionServiceOnce.Do(func() {
			transactionServiceInstance = &transactionService{
				iTransactionRepository: transactionRepo,
				iAccountRepository:     accountRepo,
			}
		})
	}
	return transactionServiceInstance
}

func (service *transactionService) Create(ctx context.Context, item *dto.CreateTransactionDTO) (interface{}, error) {
	transaction := new(models.Transaction)
	transaction.Amount = item.Amount
	transaction.Id = uuid.New()
	transaction.Unit = "INR"
	transaction.Status = "CREATED"
	transaction.Type = item.Type
	transaction.CreatedBy = ctx.Value("user.id").(string)
	destinationId, err := uuid.Parse(item.Destination)
	if err != nil {
		log.Error("Destination id incorrect")
		return nil, err
	}
	destinationAccount, err := service.iAccountRepository.Find(ctx, destinationId)
	if err != nil {
		log.Error("Failed to fetch destination account")
		return nil, err
	}
	if len(*destinationAccount) == 0 {
		log.Error("Failed to fetch destination account")
		return nil, errors.New("destination account not present")
	}
	destinationAccount0 := (*destinationAccount)[0]
	transaction.DestinationId = destinationId

	if item.Source == "" {
		transaction.SourceID = uuid.Nil
	} else {
		sourceId, err := uuid.Parse(item.Source)
		if err != nil {
			log.Error("Source account id incorrect")
			return nil, err
		}
		transaction.SourceID = sourceId
		sourceAccount, err := service.iAccountRepository.Find(ctx, sourceId)
		if err != nil {
			log.Error("Failed to fetch source account")
			return nil, err
		}
		if len(*sourceAccount) == 0 {
			log.Error("Failed to fetch source account")
			return nil, errors.New("source account not available")
		}
		sourceAccount0 := (*sourceAccount)[0]

		if sourceAccount0.Balance < transaction.Amount {
			log.Error("Insufficient balance in source account")
			return nil, errors.New("insufficient balance in source account")
		}

		err = service.iAccountRepository.UpdateBalance(ctx, sourceId, sourceAccount0.Balance-transaction.Amount)
		if err != nil {
			log.Error("Transaction Failed")
			return nil, err
		}

	}

	err = service.iAccountRepository.UpdateBalance(ctx, destinationId, destinationAccount0.Balance+transaction.Amount)
	if err != nil {
		return nil, err
	}
	trans, err := service.iTransactionRepository.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return trans, nil
}

func (service *transactionService) AddInterest(ctx context.Context, id uuid.UUID) error {
	destinationAccount, err := service.iAccountRepository.Find(ctx, id)
	if err != nil {
		log.Error("Failed to fetch destination account")
		return err
	}
	if len(*destinationAccount) == 0 {
		log.Error("Failed to fetch destination account")
		return errors.New("destination account not present")
	}
	destinationAccount0 := (*destinationAccount)[0]
	interest := uint(float64(destinationAccount0.Balance) * RATE)

	requestBody := new(dto.CreateTransactionDTO)
	requestBody.Source = ""
	requestBody.Destination = id.String()
	requestBody.Amount = interest
	requestBody.Type = "INTEREST"
	_, err = service.Create(ctx, requestBody)
	if err != nil {
		log.Error("Failed to add interest")
		return err
	}
	return nil
}

func (service *transactionService) UpdateStatus(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (service *transactionService) GetTransaction(ctx context.Context, accountId uuid.UUID, from time.Time, to time.Time) (string, error) {
	transactions, err := service.iTransactionRepository.GetTransaction(ctx, accountId, to, from)
	if err != nil {
		return "", err
	}
	transactionReport := new(dto.TransactionReport)
	transactionReport.AccountID = accountId.String()
	transactionReport.From = from
	transactionReport.To = to
	var transactionsList []dto.Transaction
	for _, item := range *transactions {
		transaction := new(dto.Transaction)
		if item.SourceID == accountId {
			transaction.AccountID = item.DestinationId.String()
			transaction.Type = "DEBIT"
		} else {
			transaction.AccountID = item.SourceID.String()
			transaction.Type = "CREDIT"
		}
		if transaction.AccountID == uuid.Nil.String() {
			transaction.AccountID = "CASH"
		}
		transaction.Time = item.CreatedOn
		transaction.Amount = covertPaisaToRupeeString(item.Amount)
		transaction.TransactionID = item.Id.String()

		transactionsList = append(transactionsList, *transaction)
	}
	transactionReport.Transactions = transactionsList
	filePath := fmt.Sprintf("%s-%s-%s.pdf", accountId, from, to)
	r := NewRequestPdf("")
	if err := r.ParseTemplate("internal/template/transactionTemplate.html", transactionReport); err == nil {
		_, err = r.GeneratePDF(filePath)
		if err != nil {
			fmt.Println("Failed to create pdf. Error: ", err)
			return "", errors.New("failed to create pdf")
		}
		return filePath, err
	} else {
		fmt.Println("Failed to create pdf. Error: ", err)
		return "", errors.New("failed to create pdf")
	}

}

func covertPaisaToRupeeString(amountInPaisa uint) string {
	inr := fmt.Sprint(amountInPaisa)
	n := len(inr)
	if n == 1 {
		return "0.0" + inr
	} else if n == 2 {
		return "0." + inr
	}
	return inr[:n-2] + "." + inr[n-2:]
}
