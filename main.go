package main

import (
	"os"
	"ukprakerja/configs"
	"ukprakerja/routes"
)

func main() {
	configs.ConnectDatabase()
	e := routes.InitRoute()

	// Menghubungkan ke PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":" + port)
}
