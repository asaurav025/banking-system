package services

import (
	"banking-system/internal/dto"
	"banking-system/internal/interfaces/irepositories"
	"banking-system/internal/models"
	"context"
	"sync"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type customerService struct {
	ICustomerRepository   irepositories.ICustomerRepository
	IKycDetailsRepository irepositories.IKycDetailsRepository
}

// For Singleton
var customerServiceOnce sync.Once
var customerServiceInstance *customerService

func NewCustomerService(
	customerRepo irepositories.ICustomerRepository,
	kycDetailsRepo irepositories.IKycDetailsRepository,
) *customerService {
	if customerServiceInstance == nil {
		customerServiceOnce.Do(func() {
			customerServiceInstance = &customerService{
				ICustomerRepository:   customerRepo,
				IKycDetailsRepository: kycDetailsRepo,
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
	customer := (*customers)[0]
	kycDetails, err := service.IKycDetailsRepository.Get(ctx, customer.KycDetailsId)
	if err != nil {
		log.Error("Failed to fetch KYC details: ", err.Error())
		return dto.GetCustomerDetailsDTO{}, err
	}
	kycDetails0 := (*kycDetails)[0]
	response.Id = customer.Id
	response.Name = customer.Name
	// response.AccountDetails = customer.AccountDetails
	response.KycDetails.GovtIDNumber = kycDetails0.GovtIdNumber
	response.KycDetails.Status = kycDetails0.Status
	response.KycDetails.Id = kycDetails0.Id
	response.KycDetails.ExpiryDate = kycDetails0.ExpiryDate.String()
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
	// To be impleneted
	return nil
}
func (service *customerService) RemoveAccount(ctx context.Context, id uuid.UUID, accountId uuid.UUID) error {
	// To be impleneted
	return nil
}
func (service *customerService) UpdateKyc(ctx context.Context, id uuid.UUID, kycId uuid.UUID) error {
	items, err := service.ICustomerRepository.Find(ctx, id)
	if err != nil {
		return err
	}
	item0 := (*items)[0]
	item0.KycDetailsId = kycId

	_, err = service.ICustomerRepository.Update(ctx, &item0)
	if err != nil {
		return err
	}
	return nil
}
