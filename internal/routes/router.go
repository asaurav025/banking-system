package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

type Route struct {
	Router *echo.Echo
}

func (r *Route) Init() {

	employeeHandler, customerHandler := Factory().InjectDepencies()

	applicationGroup := r.Router.Group("/banking-system")
	{
		applicationGroup.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "I am healthy")
		})

		v1 := applicationGroup.Group("/v1")
		{
			v1.POST("/employee/add", employeeHandler.AddEmployee)
			v1.DELETE("/employee/:id", employeeHandler.DeleteEmployee)

			v1.POST("/customer/add", customerHandler.AddCustomer)
			v1.DELETE("/customer/:id", customerHandler.DeleteCustomer)
			v1.GET("/customer/:id", customerHandler.GetCustomer)
			v1.POST("/add/kyc/customer/:customerId", customerHandler.UpdateKyc)
		}
	}

}
