package main

import (
	"log"
	"net/http"
	"order-management/models"
	"order-management/utils"
	"os"

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
}
