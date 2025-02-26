package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `gorm:"unique"`
	Age   int    `json:"age"`
}

func createUser(db *gorm.DB, user User) map[string]any {
	result := db.Create(&user)
	if result.Error != nil {
		errorjson := map[string]any{"error": result.Error.Error()}
		return errorjson
	}
	return map[string]any{"success": "Usurio Criado com sucesso!"}
}

func getUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func changeUserName(db *gorm.DB, name string, id int) error {
	var user User
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

func changeUserAge(db *gorm.DB, age int, id int) error {
	var user User
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

func deleteUser(db *gorm.DB, id int) error {
	var user User
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

func main() {

	dsn := "host=localhost user=postgres password=senha123 dbname=cinema_database port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("Erro ao criar tabela user", err)
		return
	}

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		users, err := getUsers(db)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Users": users,
		})
	})
	r.POST("/createuser", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Corpo do json mal formado",
				"details": err.Error(),
			})
			return
		}
		result := createUser(db, user)
		c.JSON(200, gin.H{
			"details": result,
		})
	})
	r.GET("/userbyid/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		result := db.First(&user, id)
		if result.Error != nil || user.ID == 0 {
			c.JSON(404, gin.H{
				"error": "Usuario não existe!!",
			})
			return
		}
		c.JSON(200, user)
	})
	r.POST("/userchangename", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Corpo do json mal formado",
				"details": err.Error(),
			})
			return
		}
		id, idExists := jsonData["id"]
		name, nameExist := jsonData["name"]
		nameStr, ok := name.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Nome não é uma string!",
			})
			return
		}
		idFloat, ok := id.(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id não é um numero!",
			})
			return
		}
		idInt := int(idFloat)
		if !idExists || !nameExist {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Esta faltando name ou id",
			})
			return
		}
		result := changeUserName(db, nameStr, idInt)
		if result != nil {
			c.JSON(500, gin.H{
				"error": result.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
			"error ": result,
		})
	})
	r.POST("/userchangeage", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Corpo do json mal formado",
				"details": err.Error(),
			})
			return
		}
		id, idExists := jsonData["id"]
		age, ageExist := jsonData["age"]
		ageFloat, ok := age.(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Age não é um numero!",
			})
			return
		}
		idFloat, ok := id.(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id não é um numero!",
			})
			return
		}
		idInt := int(idFloat)
		ageInt := int(ageFloat)
		if !idExists || !ageExist {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Esta faltando age ou id",
			})
			return
		}
		result := changeUserAge(db, ageInt, idInt)
		if result != nil {
			c.JSON(500, gin.H{
				"error": result.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "success",
			"error ": result,
		})
	})
	r.GET("/userdelete/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		result := deleteUser(db, id)
		if result != nil {
			c.JSON(500, gin.H{
				"error": result.Error(),
			})
			return
		}
	})

	r.Run(":8080")
}
