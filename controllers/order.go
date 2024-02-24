package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrder(c echo.Context) error {
	order := models.Orders{}
	c.Bind(&order)

	if err := models.DB.Create(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, order)
}

func ListOrders(c echo.Context) error {
	order := []models.Orders{}

	if err := models.DB.Find(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, order)
}

func GetOrder(c echo.Context) error {
	order := models.Orders{}

	orderIdString := c.Param("id")
	orderId, _ := strconv.Atoi(orderIdString)

	if err := models.DB.First(&order, orderId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "order not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, order)
}
func DeleteOrder(c echo.Context) error {
	order := models.Orders{}

	orderIdString := c.Param("id")
	orderId, _ := strconv.Atoi(orderIdString)

	if err := models.DB.First(&order, orderId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "order not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	if err := models.DB.Delete(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order delete sucessfully",
	})
}

func CreateOrderItem(c echo.Context) error {
	orderItem := models.OrderItems{}
	c.Bind(&orderItem)

	if err := models.DB.Create(&orderItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, orderItem)

}

func ListOrderItems(c echo.Context) error {
	orderItem := []models.OrderItems{}

	if err := models.DB.Find(&orderItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, orderItem)

}

func GetOrderItem(c echo.Context) error {
	orderItem := models.OrderItems{}

	orderItemIdstring := c.Param(":orderItemId")
	orderItemId, _ := strconv.Atoi(orderItemIdstring)

	if err := models.DB.First(&orderItem, orderItemId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "order item not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, orderItem)
}

func DeleteOrderItem(c echo.Context) error {
	orderItem := models.OrderItems{}

	orderItemIdString := c.Param(":orderItemId")
	orderItemId, _ := strconv.Atoi(orderItemIdString)

	if err := models.DB.First(&orderItem, orderItemId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "orderItem not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	if err := models.DB.Delete(&orderItem, orderItemId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete order item sucessfully",
	})
}
