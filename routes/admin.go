package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Group) {
	adminRoute := e.Group("/admins")

	adminRoute.POST("", controllers.CreateAdmin)
	adminRoute.POST("/type", controllers.CreateAdminType)
	adminRoute.GET("/:id", controllers.GetAdmin)
	adminRoute.GET("", controllers.ListAdmin)
	adminRoute.GET("/types", controllers.ListAdminTypes)
	adminRoute.DELETE("/:id", controllers.DeleteAdminById)
	adminRoute.DELETE("/type/:typeId", controllers.DeleteAdminTypeById)
}
