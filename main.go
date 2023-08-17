package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	userRepo "service-api/src/main/app/users/repository/impl"

	authJWTService "service-api/src/main/adapter/auth"
	userService "service-api/src/main/app/users/service/impl"

	userController "service-api/src/main/adapter/api/users"

	infra "service-api/src/main/infra"
)

func main() {
	fmt.Println("Hello, World!")
	db := infra.InitDB()
	infra.AutoMigrateDB(db)

	repositoryUser := userRepo.NewUserRepository(db)

	jwtService := authJWTService.NewJwtService()
	serviceUser := userService.NewUserService(repositoryUser)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		userController.UserRoutes(api, serviceUser, jwtService)
	}
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
