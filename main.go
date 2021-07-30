package main

import (
	"banking-system/internal/db"
	"banking-system/internal/routes"
	"banking-system/pkg/config"

	"github.com/labstack/echo"
)

func main() {
	config.Init()
	db.Init()
	e := echo.New()
	routes := &routes.Route{
		Router: e,
	}
	routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
