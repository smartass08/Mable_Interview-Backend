package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"

	v1 "demo/api/v1"
	"demo/config"
	"demo/logging"
	"demo/service"
)

func init() {
	logLevel, err := zapcore.ParseLevel("")
	if err != nil {
		log.Fatalln(err)
	}
	logging.Init(logLevel)
}

func main() {
	cfg := config.Get()

	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(
		fiber.Config{
			Prefork: cfg.Prefork, ErrorHandler: service.HandleError, BodyLimit: cfg.RequestBodyLimit,
		},
	)

	// Create a /api/v1 endpoint
	v1endpoint := app.Group("/api/v1")

	// Initialize controller
	controller := v1.Controller{}

	// Bind handlers
	v1.AddRoutes(v1endpoint, controller)

	// Listen on port defined in .env file
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Port)))
}
