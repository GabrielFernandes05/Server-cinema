package routes

import (
	"cinema/controllers"
	"cinema/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/getusers", func(c *gin.Context) {
			users, err := controllers.GetUsers(db)
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
		userGroup.POST("/createuser", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Corpo do json mal formado",
					"details": err.Error(),
				})
				return
			}
			result := controllers.CreateUser(db, user)
			c.JSON(200, gin.H{
				"details": result,
			})
		})
		userGroup.GET("/userbyid/:id", func(c *gin.Context) {
			var user models.User
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
		userGroup.POST("/userchangename", func(c *gin.Context) {
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
			result := controllers.ChangeUserName(db, nameStr, idInt)
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
		userGroup.POST("/userchangeage", func(c *gin.Context) {
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
			result := controllers.ChangeUserAge(db, ageInt, idInt)
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
		userGroup.GET("/userdelete/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err,
				})
				return
			}
			result := controllers.DeleteUser(db, id)
			if result != nil {
				c.JSON(500, gin.H{
					"error": result.Error(),
				})
				return
			}
		})
	}
}
