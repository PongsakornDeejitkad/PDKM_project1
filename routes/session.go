package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func SessionRoutes(e *echo.Group) {
	sessionRoutes := e.Group("/sessions")

	sessionRoutes.GET("", controllers.ListSessions)
	sessionRoutes.GET("/cartItems", controllers.ListCartItems)
	sessionRoutes.GET("/cartItem/:id", controllers.GetCartItems)
	sessionRoutes.POST("", controllers.CreateSession)
	sessionRoutes.POST("/cartItem", controllers.CreateCartItems)
	sessionRoutes.DELETE("/cartItem/:id", controllers.DeleteCartItems)

}
