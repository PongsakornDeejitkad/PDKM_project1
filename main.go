package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"order-management/models"
	"order-management/routes"
	"order-management/utils"
	"os"
	"time"
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

	v1 := e.Group("/v1")
	// v1.Use(utils.JWTMiddleware) waiting for middleware

	routes.CustomerRoutes(v1)

	// write routes details to routes.json
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return
	}
	os.WriteFile("routes.json", data, 0644)

	serveGracefulShutdown(e)
}

func serveGracefulShutdown(e *echo.Echo) {
	go func() {
		var port string
		port = os.Getenv("HTTP_PORT")
		if port == "" {
			port = utils.ViperGetString("http.port")
		}

		if err := e.Start(port); err != nil {
			log.Println("shutting down the server", err)

		}
	}()

	gracefulShutdownTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
