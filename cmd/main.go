package main

import (
	"fmt"
	"os"

	"github.com/gustavolimam/coinext-challenge/internal/environment"
	"github.com/gustavolimam/coinext-challenge/internal/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	godotenv.Load()

	environment.CheckEnvVars()

	log.Info().Msg("Starting the application...")

	e := loadServer()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv(environment.Port))))
}

func loadServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	handler := handler.New()

	e.POST("/user", handler.CreateUser)
	e.PUT("/user/:id/item", handler.AddOrRemoveItem)
	e.POST("/user/:id/trade", handler.Trade)

	return e
}
