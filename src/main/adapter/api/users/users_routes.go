package users

import (
	serviceJWT "service-api/src/main/adapter/auth"
	serviceUser "service-api/src/main/app/users/service"
	infra "service-api/src/main/infra"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.RouterGroup, service serviceUser.UserService, jwt serviceJWT.Service) {
	controller := NewUserController(service, jwt)

	userRoute := routes.Group("/users")
	{
		userRoute.GET("/", infra.AuthMiddleware(jwt, service), controller.GetAllUser)
		userRoute.GET("/:id", infra.AuthMiddleware(jwt, service), controller.GetUserByID)
		userRoute.POST("/register", controller.RegisterUser)
		userRoute.POST("/login", controller.LoginUser)
	}
}
