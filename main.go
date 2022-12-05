package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	controller "github.com/Kourin1996/go-api-e2e-sample/controllers/books"
	repository "github.com/Kourin1996/go-api-e2e-sample/repositories/books"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	ENV_FILE     = ".env"
	ENV_API_PORT = "API_PORT"
)

type Config struct {
	Port uint
}

func loadConfig() (*Config, error) {
	var err error

	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	rawPort := os.Getenv(ENV_API_PORT)
	if rawPort == "" {
		rawPort = "8000"
	}

	port, err := strconv.ParseUint(rawPort, 10, 32)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: uint(port),
	}, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)

		return
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	booksRepo := repository.NewRepository()

	controller.NewController(e, booksRepo)

	// Start server
	if err := e.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
		e.Logger.Fatal(err)
	}
}
