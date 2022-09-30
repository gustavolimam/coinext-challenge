package main

import (
	"fmt"
	"os"

	"github.com/gustavolimam/coinext-challenge/internal/environment"
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

	e.POST("/user", nil)
	e.PUT("/user/:id/item", nil)
	e.POST("/user/:id/trade", nil)

	return e
}
