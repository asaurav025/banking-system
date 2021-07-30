package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

type Route struct {
	Router *echo.Echo
}

func (r *Route) Init() {

	employeeHandler := Factory().InjectDepencies()

	applicationGroup := r.Router.Group("/banking-system")
	{
		applicationGroup.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "I am healthy")
		})

		v1 := applicationGroup.Group("/v1")
		{
			v1.POST("/employee/add", employeeHandler.AddEmployee)
			v1.DELETE("/employee/:id", employeeHandler.DeleteEmployee)
		}
	}

}
