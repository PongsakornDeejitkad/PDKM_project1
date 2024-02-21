package controllers

import (
	"fmt"
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

func CreateCustomerAddress(c echo.Context) error {
	address := models.CustomerAddress{}
	c.Bind(&address)

	customerID := c.Param("customerID")
	customer := models.Customers{}
	if err := models.DB.First(&customer, customerID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Customer not found",
		})
	}
	address.CustomerID = customer.ID

	if err := models.DB.Create(&address).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	fmt.Println(address)
	return c.JSON(http.StatusOK, address)
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

func DeleteCustomer(c echo.Context) error {
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
	if err := models.DB.Delete(&customer).Error; err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete customer",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Customer deleted successfully",
	})
}
