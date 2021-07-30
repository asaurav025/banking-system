package routes

import (
	"sync"

	"banking-system/internal/db"
	"banking-system/internal/handlers"
	"banking-system/internal/repositories"
	"banking-system/internal/services"
)

type controller struct{}

func (c *controller) InjectDepencies() (*handlers.EmployeeHandler, *handlers.CustomerHandler) {
	dbInstance := db.GetDB()
	employeeRepo := repositories.NewEmployeeRepository(dbInstance)
	customerRepo := repositories.NewCustomerRepository(dbInstance)
	kycDetailsRepo := repositories.NewKycDetailsRepository(dbInstance)

	employeeService := services.NewEmployeeService(employeeRepo)
	customerService := services.NewCustomerService(customerRepo, kycDetailsRepo)
	kycDetailsService := services.NewKycDetialsService(kycDetailsRepo)

	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	customerHandler := handlers.NewCustomerHandler(customerService, kycDetailsService)

	return employeeHandler, customerHandler
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
