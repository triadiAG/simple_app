package controller

import (
	"net/http"
	"strconv"

	"myApp/config"
	"myApp/model"

	m "myApp/middleware"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var user []model.User

	if err := config.DB.Find(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

// Get one user
func GetUserController(c echo.Context) error {
	var user []model.User
	// Get ID
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Take(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get user",
		"data":    user,
	})
}

// Create Buku
func CreateUsersController(c echo.Context) error {
	// Binding data
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"data":    user,
	})
}

// Update Buku
func UpdateUserController(c echo.Context) error {
	// Binding data
	data_user := model.User{}
	c.Bind(&data_user)

	// Get ID
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Model(&model.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"username": data_user.Username,
		"email":    data_user.Email,
		"password": data_user.Password,
	}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update user",
		"data":    data_user,
	})
}

// Delete user
func DeleteUserController(c echo.Context) error {
	// Get ID
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&model.User{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user Deleted",
	})
}

func LoginController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := m.CreateToken(int(user.ID), user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{int(user.ID), user.Username, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login !",
		"user":    userResponse,
	})
}
