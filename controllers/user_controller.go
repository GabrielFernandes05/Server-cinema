package controllers

import (
	"cinema/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.User) map[string]any {
	result := db.Create(&user)
	if result.Error != nil {
		errorjson := map[string]any{"error": result.Error.Error()}
		return errorjson
	}
	return map[string]any{"success": "Usurio Criado com sucesso!"}
}

func GetUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func ChangeUserName(db *gorm.DB, name string, id int) error {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.Model(&user).Update("Name", name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ChangeUserAge(db *gorm.DB, age int, id int) error {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.Model(&user).Update("Age", age)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(db *gorm.DB, id int) error {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
