package controllers

import (
	"net/http"
	"strconv"

	"ukprakerja/models"

	"ukprakerja/configs"

	"github.com/labstack/echo/v4"
)

// Memperoleh item di dalam API
func GetItemsController(c echo.Context) error {
	var items []models.Item
	// Mencari tabel item di DB
	result := configs.DB.Find(&items)

	// Bila gagal memperoleh data di DB
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal memperoleh data", Data: nil,
		})
	}
	// Menampilkan data
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil", Data: items,
	})
}

// Memasukkan item ke dalam API
func InsertItemController(c echo.Context) error {
	// Binding JSON ke dalam bentuk struct
	var insertItem models.Item
	c.Bind(&insertItem)

	// Memasukkan item baru
	result := configs.DB.Create(&insertItem)

	// Kurang validasi data yang dimasukkan

	// Bila gagal memasukkan item
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal menambah pegawai", Data: nil,
		})
	}

	// Status berhasil memasukkan item
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil menambah pegawai", Data: insertItem,
	})
}

// Memperoleh item di dalam API
func GetSingleItemController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// Kurang validasi parameter id

	// Mencari data
	var item models.Item
	result := configs.DB.First(&item, id)

	// Bila data tidak ditemukan
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Gagal memperoleh data", Data: nil,
		})
	}

	// Data ditemukan dan ditampilkan di response
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil", Data: item,
	})
}

// Melakukan update item ke DB
func UpdateItemController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// Kurang validasi parameter id

	// Binding data dari JSON ke struct
	var updateItem models.Item
	c.Bind(&updateItem)

	// Kurang validasi bila data yang disimpan salah

	// Menyimpan data item
	// Menggunakan DB model untuk 'query' where=id param berisi data updateItem
	result := configs.DB.Model(&models.Item{}).Where("id = ?", id).Updates(updateItem)

	// Bila penyimpanan gagal di server
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal memperbaharui item", Data: nil,
		})
	}

	// Status bila data berhasil disimpan
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil memperbaharui item", Data: updateItem,
	})
}

// Menghapus data item
func DeleteItemController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// Kurang validasi parameter id

	// Menghapus data id
	var item models.Item
	result := configs.DB.Delete(&item, id)

	// Bila ada data terhapus
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusOK, models.BaseResponse{
			Status: true, Message: "Berhasil menghapus item", Data: nil,
		})
	}

	// bila gagal hapus dari db
	return c.JSON(http.StatusNotFound, models.BaseResponse{
		Status: false, Message: "Item tidak ditemukan", Data: nil,
	})
}
