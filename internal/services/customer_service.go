package services

import (
	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type customerService struct {
	ICustomerRepository   irepositories.ICustomerRepository
	IKycDetailsRepository irepositories.IKycDetailsRepository
	IAccountRepository    irepositories.IAccountRepository
}

// For Singleton
var customerServiceOnce sync.Once
var customerServiceInstance *customerService

func NewCustomerService(
	customerRepo irepositories.ICustomerRepository,
	kycDetailsRepo irepositories.IKycDetailsRepository,
	accountRepo irepositories.IAccountRepository,
) *customerService {
	if customerServiceInstance == nil {
		customerServiceOnce.Do(func() {
			customerServiceInstance = &customerService{
				ICustomerRepository:   customerRepo,
				IKycDetailsRepository: kycDetailsRepo,
				IAccountRepository:    accountRepo,
			}
		})
	}
	return customerServiceInstance
}

func (service *customerService) Create(ctx context.Context, item *dto.CreateCustomerDTO) (interface{}, error) {
	customer := new(models.Customer)
	customer.Id = uuid.New()
	customer.Name = item.Name
	customer.CreatedBy = ctx.Value("user.id").(string)
	customer.AccountDetails = "{}"
	cust, err := service.ICustomerRepository.Create(ctx, customer)
	if err != nil {
		return nil, err
	}
	return cust, nil
}

func (service *customerService) GetCustomer(ctx context.Context, id uuid.UUID) (dto.GetCustomerDetailsDTO, error) {
	response := new(dto.GetCustomerDetailsDTO)
	customers, err := service.ICustomerRepository.Find(ctx, id)
	if err != nil {
		log.Error("Failed to fetch customer details: ", err.Error())
		return dto.GetCustomerDetailsDTO{}, err
	}
	if len(*customers) == 0 {
		log.Error("Failed to fetch customer details")
		return dto.GetCustomerDetailsDTO{}, errors.New("failed to fetch customer details")
	}
	customer := (*customers)[0]

	var account models.CustomerAccount
	err = json.Unmarshal([]byte(customer.AccountDetails), &account)
	if err != nil {
		return dto.GetCustomerDetailsDTO{}, err
	}
	accountList := account.Accounts
	var accountDetailsList []dto.AccountDetails
	// todo: make it asynchronous
	for _, accountId := range accountList {
		id, _ := uuid.Parse(accountId)
		val, err := service.IAccountRepository.Find(ctx, id)
		if err != nil {
			continue
		}
		if len(*val) == 0 {
			continue
		}
		accountDetails := new(dto.AccountDetails)
		val0 := (*val)[0]
		accountDetails.Balance = fmt.Sprint(val0.Balance)
		accountDetails.Type = val0.Type
		accountDetails.Id = val0.Id

		accountDetailsList = append(accountDetailsList, *accountDetails)
	}
	if customer.KycDetailsId != uuid.Nil {
		kycDetails, err := service.IKycDetailsRepository.Get(ctx, customer.KycDetailsId)
		if err != nil {
			log.Error("Failed to fetch KYC details: ", err.Error())
			return dto.GetCustomerDetailsDTO{}, err
		}
		kycDetails0 := (*kycDetails)[0]

		response.KycDetails.GovtIDNumber = kycDetails0.GovtIdNumber
		response.KycDetails.Status = kycDetails0.Status
		response.KycDetails.Id = kycDetails0.Id
		response.KycDetails.ExpiryDate = kycDetails0.ExpiryDate.String()
	}

	response.Id = customer.Id
	response.Name = customer.Name
	response.AccountDetails = accountDetailsList
	return *response, nil
}
func (service *customerService) UpdateName(ctx context.Context, name string) error {
	// To be impleneted
	return nil
}
func (service *customerService) Deactivate(ctx context.Context, id uuid.UUID) error {
	service.ICustomerRepository.Delete(ctx, id)
	return nil
}

func (service *customerService) AddAccount(ctx context.Context, id uuid.UUID, accountId uuid.UUID) error {
	customers, err := service.ICustomerRepository.Find(ctx, id)
	if err != nil {
		log.Error("Failed to fetch customer details: ", err.Error())
		return err
	}
	customer := (*customers)[0]
	var account models.CustomerAccount
	err = json.Unmarshal([]byte(customer.AccountDetails), &account)
	if err != nil {
		return err
	}
	accountList := account.Accounts
	accountList = append(accountList, accountId.String())
	accountList = removeDuplicateValues(accountList)
	temp, err := json.Marshal(&models.CustomerAccount{
		Accounts: accountList,
	})
	if err != nil {
		return err
	}
	customer.AccountDetails = string(temp)
	customer.UpdatedBy = ctx.Value("user.id").(string)
	customer.UpdatedOn = time.Now()
	_, err = service.ICustomerRepository.Update(ctx, &customer)
	if err != nil {
		log.Error("Failed to get customer details")
		return err
	}
	return nil
}

func (service *customerService) RemoveAccount(ctx context.Context, id uuid.UUID, accountId uuid.UUID) error {
	customers, err := service.ICustomerRepository.Find(ctx, id)
	if err != nil {
		log.Error("Failed to fetch customer details: ", err.Error())
		return err
	}
	customer := (*customers)[0]
	var account models.CustomerAccount
	err = json.Unmarshal([]byte(customer.AccountDetails), &account)
	if err != nil {
		return err
	}
	accountList := account.Accounts
	list := []string{}

	for _, entry := range accountList {
		if entry != accountId.String() {
			list = append(list, entry)
		}
	}
	temp, err := json.Marshal(&models.CustomerAccount{
		Accounts: list,
	})
	if err != nil {
		return err
	}
	customer.AccountDetails = string(temp)
	_, err = service.ICustomerRepository.Update(ctx, &customer)
	if err != nil {
		return err
	}
	return nil
}
func (service *customerService) UpdateKyc(ctx context.Context, id uuid.UUID, kycId uuid.UUID) error {
	items, err := service.ICustomerRepository.Find(ctx, id)
	if err != nil {
		return err
	}
	item0 := (*items)[0]
	item0.KycDetailsId = kycId
	item0.UpdatedBy = ctx.Value("user.id").(string)
	item0.UpdatedOn = time.Now()
	_, err = service.ICustomerRepository.Update(ctx, &item0)
	if err != nil {
		log.Error("Failed to update customer")
		return err
	}
	return nil
}

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
