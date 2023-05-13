package models

// Struct pegawai
type Employee struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Title string `json:"title"`
}
