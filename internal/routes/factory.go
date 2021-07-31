package routes

import (
	"sync"

	"banking-system/internal/db"
	"banking-system/internal/handlers"
	"banking-system/internal/repositories"
	"banking-system/internal/services"
)

type controller struct{}

func (c *controller) InjectDepencies() (*handlers.EmployeeHandler, *handlers.CustomerHandler, *handlers.AccountHandler) {
	dbInstance := db.GetDB()
	employeeRepo := repositories.NewEmployeeRepository(dbInstance)
	customerRepo := repositories.NewCustomerRepository(dbInstance)
	kycDetailsRepo := repositories.NewKycDetailsRepository(dbInstance)
	accountRepo := repositories.NewAccountRepository(dbInstance)

	employeeService := services.NewEmployeeService(employeeRepo)
	customerService := services.NewCustomerService(customerRepo, kycDetailsRepo, accountRepo)
	kycDetailsService := services.NewKycDetialsService(kycDetailsRepo)
	accountService := services.NewAccountService(accountRepo)

	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	customerHandler := handlers.NewCustomerHandler(customerService, kycDetailsService)
	accountHandler := handlers.NewAccountHandler(accountService)

	return employeeHandler, customerHandler, accountHandler
}

var controllerInstance *controller
var controllerOnce sync.Once

func Factory() *controller {
	if controllerInstance == nil {
		controllerOnce.Do(func() {
			controllerInstance = &controller{}
		})
	}
	return controllerInstance
}
