package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	orderRoute := e.Group("/orders")

	orderRoute.GET("", controllers.ListOrders)
	orderRoute.GET("/:id", controllers.GetOrder)
	orderRoute.POST("/:id", controllers.CreateOrder)
	orderRoute.DELETE("/:id", controllers.DeleteOrder)

}
