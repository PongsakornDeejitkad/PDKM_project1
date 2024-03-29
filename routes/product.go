package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRoute := e.Group("/products")

	productRoute.POST("", controllers.CreateProduct)
	productRoute.GET("/:id", controllers.GetProduct)
	productRoute.GET("", controllers.ListProduct)

}
