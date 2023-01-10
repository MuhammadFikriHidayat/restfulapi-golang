package controller

import (
	"apigo/config"
	model "apigo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllAdmin(c echo.Context) error {
	admin := model.Admin{}

	c.Bind(&admin)
	err := config.DB.Find(&admin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    admin,
	})
}
