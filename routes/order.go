package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	orderRoute := e.Group("/orders")

	orderRoute.GET("", controllers.ListOrders)
	orderRoute.GET("/:id", controllers.GetOrder)
	orderRoute.GET("orderItems", controllers.ListOrderItems)
	orderRoute.GET("orderItems/:orderItemId", controllers.GetOrderItem)
	orderRoute.POST("orderItems", controllers.CreateOrderItem)
	orderRoute.POST("", controllers.CreateOrder)
	orderRoute.DELETE("/:id", controllers.DeleteOrder)
	orderRoute.DELETE("orderItems/:orderItemId", controllers.DeleteOrderItem)

}
