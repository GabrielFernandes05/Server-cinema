package main

import (
	"cinema/config"
	"cinema/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.InitDB()
	defer config.CloseDB(db)

	r := gin.Default()

	routes.SetupUserRoutes(r, db)

	r.Run(":8080")
}
