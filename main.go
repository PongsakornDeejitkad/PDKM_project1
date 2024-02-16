package main

import (
	"log"
	"net/http"
	"order-management/models"
	"order-management/utils"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

var runEnv string

func init() {
	runEnv = os.Getenv("RUN_ENV")
	if runEnv == "" {
		runEnv = "dev"
	}

	utils.InitViper()

	var err error
	if err = models.ConnectDB(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))

	// Wait for an interrupt signal (e.g., Ctrl+C) to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
