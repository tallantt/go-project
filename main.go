package main

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/db"
	"rest-project/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
