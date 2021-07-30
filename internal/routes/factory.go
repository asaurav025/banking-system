package routes

import (
	"banking-system/internal/handlers"
	"banking-system/internal/repositories"
	"banking-system/internal/services"
	"banking-system/pkg/db"
	"sync"
)

type controller struct{}

func (c *controller) InjectDepencies() *handlers.EmployeeHandler {
	dbInstance := db.GetDB()

	employeeRepo := repositories.NewEmployeeRepository(dbInstance)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	return employeeHandler
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
