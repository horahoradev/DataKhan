package main

import (
	"github.com/horahoradev/DataKhan/backend/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//cfg, err := config.New()
	//if err != nil {
	//	log.Fatalf("Could not initialize config. Err: %s", err)
	//}

	e := echo.New()

	e.Use(middleware.Logger())

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
