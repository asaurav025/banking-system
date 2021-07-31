package routes

import (
	"banking-system/pkg/middleware/authentication"
	"net/http"

	"github.com/labstack/echo"
)

type Route struct {
	Router *echo.Echo
}

func (r *Route) Init() {

	employeeHandler, customerHandler, accountHandler, transactionHandler := Factory().InjectDepencies()

	applicationGroup := r.Router.Group("/banking-system")
	{
		applicationGroup.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "I am healthy")
		})

		applicationGroup.POST("/login", employeeHandler.Login)

		v1 := applicationGroup.Group("/v1", authentication.AuthenticationMiddleware())
		{

			v1.POST("/employee/add", employeeHandler.AddEmployee)
			v1.DELETE("/employee/:id", employeeHandler.DeleteEmployee)

			v1.POST("/customer/add", customerHandler.AddCustomer)
			v1.DELETE("/customer/:id", customerHandler.DeleteCustomer)
			v1.GET("/customer/:id", customerHandler.GetCustomer)
			v1.POST("/kyc/add/customer/:customerId", customerHandler.UpdateKyc)
			v1.POST("/link/customer/:customerId/account/:accountId", customerHandler.LinkAccount)

			v1.POST("/account/add", accountHandler.CreateAccount)
			v1.GET("/account/:id/balance", accountHandler.GetBalance)

			v1.POST("/transaction/create", transactionHandler.CreateTransaction)
			v1.GET("/transaction/account/:accountId/from/:from/to/:to", transactionHandler.GetTransaction)
		}
	}

}
