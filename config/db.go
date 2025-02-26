package config

import (
	"cinema/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=senha123 dbname=cinema_database port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		panic("Falha na conexão!")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Erro ao criar tabela user", err)
		panic("Erro ao criar tabelas!")
	}
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
