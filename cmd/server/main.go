package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lukiriskigumilar/resepify-be/internal/auth"
	"github.com/lukiriskigumilar/resepify-be/internal/users"

	"github.com/joho/godotenv"
	"github.com/lukiriskigumilar/resepify-be/internal/config"
	"github.com/lukiriskigumilar/resepify-be/internal/routes"
)

func main() {
	//load env
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	//init db
	config.ConnectDatabase()
	db := config.DB

	//init modules
	usersModule := users.InitUserModule(db)
	authModule := auth.InitAuthModule(usersModule)

	router := gin.Default()
	api := router.Group("/api/v1")

	//init routes
	routes.GlobalRoutes(api, authModule)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"status":  "success",
			"data":    nil,
		})
	})

	router.Run(":8080")
}
