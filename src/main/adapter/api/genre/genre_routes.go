package genre

import (
	jwt "service-api/src/main/adapter/auth"
	genreService "service-api/src/main/app/genre/service"
	userService "service-api/src/main/app/users/service"
	infra "service-api/src/main/infra"

	"github.com/gin-gonic/gin"
)

func GenreRoutes(routes *gin.RouterGroup, service genreService.GenreService, serviceUser userService.UserService, serviceAuth jwt.Service) {
	var controller = NewGenreController(service)

	genreRoute := routes.Group("/genres")
	{
		genreRoute.GET("/", infra.AuthMiddleware(serviceAuth, serviceUser), controller.GetAllGenre)
		genreRoute.GET("/:id", infra.AuthMiddleware(serviceAuth, serviceUser), controller.GetGenreByID)
		genreRoute.POST("/", infra.AuthMiddleware(serviceAuth, serviceUser), controller.CreateGenre)
		genreRoute.PUT("/:id", infra.AuthMiddleware(serviceAuth, serviceUser), controller.UpdateGenre)
		genreRoute.DELETE("/:id", infra.AuthMiddleware(serviceAuth, serviceUser), controller.DeleteGenre)
	}
}
