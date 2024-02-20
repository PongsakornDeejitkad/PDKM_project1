package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Group) {
	adminRoute := e.Group("/admins")

	adminRoute.POST("", controllers.CreateAdmin)
	adminRoute.GET("/:id", controllers.GetAdmin)
	adminRoute.GET("", controllers.ListAdmin)
}
