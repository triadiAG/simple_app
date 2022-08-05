package controller

import (
	"net/http"
	"strconv"

	"myApp/config"
	"myApp/model"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var buku []model.Books

	if err := config.DB.Find(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    buku,
	})
}

// Get one user
func GetBukuController(c echo.Context) error {
	var buku []model.Books
	// Get ID
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Take(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get buku",
		"data":    buku,
	})
}

// Create Buku
func CreateBooksController(c echo.Context) error {
	// Binding data
	buku := model.Books{}
	c.Bind(&buku)

	if err := config.DB.Save(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create buku",
		"data":    buku,
	})
}

// Update Buku
func UpdateBukuController(c echo.Context) error {
	// Binding data
	data_buku := model.Books{}
	c.Bind(&data_buku)

	// Get ID
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Model(&model.Books{}).Where("id = ?", id).Updates(map[string]interface{}{
		"judul":        data_buku.Judul,
		"pengarang":    data_buku.Pengarang,
		"penerbit":     data_buku.Penerbit,
		"tahun_terbit": data_buku.TahunTerbit,
	}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update Buku",
		"data":    data_buku,
	})
}

// Delete Buku
func DeleteBukuController(c echo.Context) error {
	// Get ID
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&model.Books{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Buku Deleted",
	})
}
