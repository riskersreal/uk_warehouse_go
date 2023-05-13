package configs

import (
	"fmt"
	"os"
	"ukprakerja/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // biar bisa diakses file lain. harus pointer agar value di DB berubah

// Menghubungkan REST API dengan database
func ConnectDatabase() {
	// Menyesuaikan database agar sesuai dengan API
	loadenv()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init database failed") // Panic bila tidak bisa buka database
	}
	//
	migrateDatabase()
}

// Meload file env untuk database
func loadenv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}

// Auto migrate agar database sesuai dengan file lokal
func migrateDatabase() {
	DB.AutoMigrate(&models.Item{})
	DB.AutoMigrate(&models.Employee{})
}
