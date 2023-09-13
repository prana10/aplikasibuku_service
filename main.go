package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	bookRepo "service-api/src/main/app/book/repository/impl"
	genreRepo "service-api/src/main/app/genre/repository/impl"
	userRepo "service-api/src/main/app/users/repository/impl"

	authJWTService "service-api/src/main/adapter/auth"
	bookService "service-api/src/main/app/book/service/impl"
	genreService "service-api/src/main/app/genre/service/impl"
	userService "service-api/src/main/app/users/service/impl"

	bookController "service-api/src/main/adapter/api/book"
	genreController "service-api/src/main/adapter/api/genre"
	userController "service-api/src/main/adapter/api/users"

	infra "service-api/src/main/infra"
)

func main() {
	fmt.Println("Hello, World!")
	db := infra.InitDB()
	infra.AutoMigrateDB(db)

	repositoryUser := userRepo.NewUserRepository(db)
	repositoryBook := bookRepo.NewBookRepository(db)
	repositoryGenre := genreRepo.NewGenreRepository(db)

	jwtService := authJWTService.NewJwtService()
	serviceUser := userService.NewUserService(repositoryUser)
	serviceGenre := genreService.NewGenreService(repositoryGenre)
	serviceBook := bookService.NewBookService(repositoryBook, repositoryGenre)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		userController.UserRoutes(api, serviceUser, jwtService)
		bookController.BookRoutes(api, serviceUser, serviceBook, serviceGenre, jwtService)
		genreController.GenreRoutes(api, serviceGenre, serviceUser, jwtService)
	}
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":3000")
}
