package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	if err := models.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, product)
}

func GetProduct(c echo.Context) error {
	product := models.Product{}

	productIdString := c.Param("id")
	productId, _ := strconv.Atoi(productIdString)

	if err := models.DB.First(&product, productId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Product not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, product)
}

func ListProduct(c echo.Context) error {
	products := []models.Product{}

	if err := models.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, products)
}
