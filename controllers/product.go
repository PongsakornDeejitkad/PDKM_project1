package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func CreateProduct(c echo.Context) error {
	product := models.Products{}
	c.Bind(&product)

	if err := models.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, product)
}

func GetProduct(c echo.Context) error {
	product := models.Products{}

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
	products := []models.Products{}

	if err := models.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, products)
}

func DeleteProductById(c echo.Context) error {
	product := models.Products{}

	productIdString := c.Param("id")
	productId, _ := strconv.Atoi(productIdString)

	if err := models.DB.First(&product, productId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})

		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := models.DB.Delete(&product).Error; err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete customer",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Customer deleted successfully",
	})
}

func CreateProductCategory(c echo.Context) error {
	category := models.ProductsCategory{}
	c.Bind(&category)
	if err := models.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, category)
}

func DeleteProductCategoryById(c echo.Context) error {
	category := models.ProductsCategory{}

	categoryIdString := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdString)
	if err := models.DB.First(&category, categoryId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "categoryId not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := models.DB.Delete(&category).Error; err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete Product category",
		})
	}
	return c.JSON(http.StatusOK, category)

}
