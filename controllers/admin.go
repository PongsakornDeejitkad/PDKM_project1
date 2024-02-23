package controllers

import (
	"net/http"
	"order-management/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
func DeleteAdminById(c echo.Context) error {
	admin := models.Admins{}

	adminIdString := c.Param("id")
	adminId, _ := strconv.Atoi(adminIdString)
	if err := models.DB.First(&admin, adminId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "admin not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := models.DB.Delete(&admin).Error; err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete admin",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted admin successfully",
	})

}

func CreateAdminType(c echo.Context) error {
	adminType := models.AdminType{}
	c.Bind(&adminType)
	if err := models.DB.Create(&adminType).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, adminType)
}
func ListAdminTypes(c echo.Context) error {
	adminType := []models.AdminType{}
	if err := models.DB.Find(&adminType).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, adminType)
}

func DeleteAdminTypeById(c echo.Context) error {
	adminType := models.AdminType{}

	adminTypeIdString := c.Param("typeId")
	typeId, _ := strconv.Atoi(adminTypeIdString)
	if err := models.DB.First(&adminType, typeId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Admin type not Found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := models.DB.Delete(&adminType).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete admin type",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Admin type deleted successfully",
	})

}
