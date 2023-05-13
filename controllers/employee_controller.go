package controllers

import (
	"net/http"
	"strconv"

	"ukprakerja/models"

	"ukprakerja/configs"

	"github.com/labstack/echo/v4"
)

// Memperoleh pegawai di dalam API
func GetEmployeesController(c echo.Context) error {
	var employees []models.Employee
	// Mencari seluruh data pegawai
	result := configs.DB.Find(&employees)

	// Error bila tidak dapat memperoleh data
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal memperoleh data pegawai", Data: nil,
		})
	}
	// Menampilkan data
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil", Data: employees,
	})
}

// Memasukkan pegawai ke dalam API
func InsertEmployeeController(c echo.Context) error {
	// Binding pegawai dari bentuk JSON ke struct
	var insertEmployee models.Employee
	c.Bind(&insertEmployee)

	// Kurang validasi data yang dimasukkan

	// Memasukkan pegawai baru
	result := configs.DB.Create(&insertEmployee)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal menambah pegawai", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil menambah pegawai", Data: insertEmployee,
	})
}

// Memperoleh satu pegawai di dalam API
func GetSingleEmployeeController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// Kurang validasi bila data id tidak valid

	// Mencari data pegawai berdasarkan param
	var employee models.Employee
	result := configs.DB.First(&employee, id)

	// Bila data tidak ditemukan
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Gagal menemukan pegawai", Data: nil,
		})
	}

	// Data ditemukan dan dimunculkan
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil", Data: employee,
	})
}

// Melakukan update pegawai ke DB
func UpdateEmployeeController(c echo.Context) error {
	// Mengambil parameter pegawai
	id, _ := strconv.Atoi(c.Param("id"))
	// Kurang validasi bila data id tidak valid

	// Binding data input JSON ke struct
	var updateEmployee models.Employee
	c.Bind(&updateEmployee)

	// Kurang validasi bila data yang disimpan salah

	// Menyimpan update data pegawai
	// DB Model pada employee dengan 'query' where id dari Where dan update sesuai data updateEmployee
	result := configs.DB.Model(&models.Employee{}).Where("id = ?", id).Updates(updateEmployee)
	// Bila gagal menyimpan data
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal memperbaharui pegawai", Data: nil,
		})
	}
	// Bila berhasil menyimpan data
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Pegawai berhasil diperbaharui", Data: updateEmployee,
	})
}

// Hapus data pegawai dari DB
func DeleteEmployeeController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee models.Employee
	result := configs.DB.Delete(&employee, id)

	// Bila ada data terhapus
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusOK, models.BaseResponse{
			Status: true, Message: "Berhasil menghapus pegawai", Data: nil,
		})
	}

	// Bila gagal hapus
	return c.JSON(http.StatusNotFound, models.BaseResponse{
		Status: false, Message: "Pegawai tidak ditemukan", Data: nil,
	})
}
