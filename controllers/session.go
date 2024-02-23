package controllers

import (
	"net/http"
	"order-management/models"

	"github.com/labstack/echo/v4"
)

func ListSessions(c echo.Context) error {
	session := models.Sessions{}
	c.Bind(&session)

	if err := models.DB.Find(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, session)
}
