package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"routes"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Could not initialize config. Err: %s", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())

	grpcAuth := custommiddleware.NewGRPCAuth(cfg)
	e.Use(grpcAuth.GRPCAuth)

	routes.SetupRoutes(e, cfg)
}
