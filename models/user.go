package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `gorm:"unique"`
	Age   int    `json:"age"`
}
