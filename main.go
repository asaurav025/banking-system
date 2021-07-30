package main

import (
	"banking-system/internal/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes := &routes.Route{
		Router: e,
	}
	routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
