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

func GetAdminById(c echo.Context) error {
	admin := model.Admin{}

	c.Bind(&admin)
	err := config.DB.Where("id = ? ", admin.Id).First(&admin).Error 
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": admin.Id,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":admin,
	})
}

func CreateAdmin(c echo.Context) error {
	createAdmin := model.Admin{}

	c.Bind(&createAdmin)

	err := config.DB.Save(&createAdmin).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating admin",
		"data": createAdmin,
	})
}
