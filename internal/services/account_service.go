package services

import (
	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type accountService struct {
	iAccountRepository irepositories.IAccountRepository
}

// For Singleton
var accountServiceOnce sync.Once
var accountServiceInstance *accountService

func NewAccountService(
	accountRepo irepositories.IAccountRepository,
) *accountService {
	if accountServiceInstance == nil {
		accountServiceOnce.Do(func() {
			accountServiceInstance = &accountService{
				iAccountRepository: accountRepo,
			}
		})
	}
	return accountServiceInstance
}

func (service *accountService) Create(ctx context.Context, item *dto.AccountRequestDto) (*models.Account, error) {

	account := new(models.Account)
	account.Id = uuid.New()
	account.Type = item.Type
	account.Balance = 0
	account.CreatedBy = ctx.Value("user.id").(string)
	account.Unit = "INR"
	acc, err := service.iAccountRepository.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (service *accountService) GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	accs, err := service.iAccountRepository.Find(ctx, id)
	if err != nil {
		log.Error("Failed to fetch customer details: ", err.Error())
		return nil, err
	}
	if len(*accs) == 0 {
		log.Error("Failed to fetch account")
		return nil, errors.New("account not available")
	}
	acc := (*accs)[0]

	return &acc, nil
}

func (service *accountService) UpdateBalance(ctx context.Context, id uuid.UUID, balance uint) error {
	return service.iAccountRepository.UpdateBalance(ctx, id, balance)
}
