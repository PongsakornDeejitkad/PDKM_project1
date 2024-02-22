package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAdmin(c echo.Context) error {
	admin := models.Admins{}
	c.Bind(&admin)
	if err := models.DB.Create(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, admin)
}

func GetAdmin(c echo.Context) error {
	admin := models.Admins{}

	adminIdString := c.Param("id")
	adminId, _ := strconv.Atoi(adminIdString)
	if err := models.DB.First(&admin, adminId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Admin not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, admin)

}

func ListAdmin(c echo.Context) error {
	admin := []models.Admins{}

	if err := models.DB.Find(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})

	}
	return c.JSON(http.StatusOK, admin)
}
