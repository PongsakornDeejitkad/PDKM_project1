package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func SessionRoutes(e *echo.Group) {
	sessionRoutes := e.Group("/sessions")

	sessionRoutes.GET("", controllers.ListSessions)
	sessionRoutes.GET("/:id", controllers.GetSession)
	sessionRoutes.POST("", controllers.CreateSession)
	sessionRoutes.GET("/:id/cartItem", controllers.GetCartItems)
}
