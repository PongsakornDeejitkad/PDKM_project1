package routes

import (
	"order-management/controllers"

	"github.com/labstack/echo/v4"
)

func CustomerRoutes(e *echo.Group) {
	customerRoute := e.Group("/customers")

	customerRoute.POST("", controllers.CreateCustomer)
	customerRoute.POST("/:customerId/address", controllers.CreateCustomerAddress)
	customerRoute.POST("/:customerId/payment", controllers.CreateCustomerPayment)
	customerRoute.GET("/:id", controllers.GetCustomer)
	customerRoute.GET("", controllers.ListCustomer)
	customerRoute.DELETE("/:id", controllers.DeleteCustomerById)

}
