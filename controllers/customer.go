package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateCustomer(c echo.Context) error {
	customer := models.Customers{}
	c.Bind(&customer)

	if err := models.DB.Create(&customer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, customer)
}

func GetCustomer(c echo.Context) error {
	customer := models.Customers{}

	customerIdString := c.Param("id")
	customerId, _ := strconv.Atoi(customerIdString)

	if err := models.DB.First(&customer, customerId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Customer not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, customer)
}

func ListCustomer(c echo.Context) error {
	customers := []models.Customers{}

	if err := models.DB.Find(&customers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, customers)
}
