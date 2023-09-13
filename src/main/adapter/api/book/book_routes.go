package book

import (
	jwt "service-api/src/main/adapter/auth"
	serviceBook "service-api/src/main/app/book/service"
	serviceGenre "service-api/src/main/app/genre/service"
	serviceUser "service-api/src/main/app/users/service"
	infra "service-api/src/main/infra"

	"github.com/gin-gonic/gin"
)

func BookRoutes(routes *gin.RouterGroup, userService serviceUser.UserService, service serviceBook.BookService, serviceGenre serviceGenre.GenreService, jwtService jwt.Service) {
	var controller = NewBookController(service, serviceGenre)

	bookRoute := routes.Group("/books")
	{
		bookRoute.POST("/", infra.AuthMiddleware(jwtService, userService), controller.InsertBook)
		bookRoute.GET("/", controller.GetAllBook)
		bookRoute.GET("/:id", controller.GetBookByID)
		bookRoute.PUT("/:id", infra.AuthMiddleware(jwtService, userService), controller.UpdateBook)
		bookRoute.DELETE("/:id", infra.AuthMiddleware(jwtService, userService), controller.DeleteBookByID)
	}
}
