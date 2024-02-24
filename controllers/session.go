package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateSession(c echo.Context) error {
	session := models.Sessions{}
	c.Bind(&session)
	if err := models.DB.Create(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, session)
}

func ListSessions(c echo.Context) error {
	session := []models.Sessions{}
	if err := models.DB.Find(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, session)
}

func DeleteSession(c echo.Context) error {
	session := models.Sessions{}

	sessionIdString := c.Param("id")
	sessionId, _ := strconv.Atoi(sessionIdString)
	if err := models.DB.Delete(&session, sessionId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "session not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "session deleted successfully",
	})
}

func CreateCartItems(c echo.Context) error {
	cart := models.CartItems{}
	c.Bind(&cart)

	if err := models.DB.Create(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, cart)
}

func ListCartItems(c echo.Context) error {
	cart := []models.CartItems{}

	if err := models.DB.Find(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, cart)
}

func GetCartItems(c echo.Context) error {

	cart := models.CartItems{}

	cartIdString := c.Param("id")
	cartId, _ := strconv.Atoi(cartIdString)

	if err := models.DB.First(&cart, cartId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "cart not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, cart)
}

func DeleteCartItems(c echo.Context) error {
	cart := models.CartItems{}

	cartIdstring := c.Param("id")
	cartId, _ := strconv.Atoi(cartIdstring)

	if err := models.DB.Delete(&cart, cartId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "cart not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": " cart deleted successfully",
	})
}
