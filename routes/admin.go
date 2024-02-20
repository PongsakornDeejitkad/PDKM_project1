package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Group) {
	customerRoute := e.Group("/admins")

	customerRoute.POST("", controllers.CreateCustomer)
	customerRoute.GET("/:id", controllers.GetCustomer)
	customerRoute.GET("", controllers.ListCustomer)
}
