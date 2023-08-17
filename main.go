package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// authJWTService "service-api/src/main/adapter/auth"
	infra "service-api/src/main/infra"
)

func main() {
	fmt.Println("Hello, World!")
	db := infra.InitDB()
	infra.AutoMigrateDB(db)

	// jwtService := authJWTService.NewJwtService()

	router := gin.Default()
	api := router.Group("/api/v1")
	{

	}
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
