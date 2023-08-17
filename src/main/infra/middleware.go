package infra

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	authJWT "service-api/src/main/adapter/auth"
	serviceUser "service-api/src/main/app/users/service"
)

func AuthMiddleware(serviceAuth authJWT.Service, userService serviceUser.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := NewResponseAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)

			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := serviceAuth.ValidateToken(tokenString)

		if err != nil {
			response := NewResponseAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)

			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := NewResponseAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)

			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(uint(userID))
		if err != nil {
			response := NewResponseAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)

			return
		}

		ctx.Set("user", user)
		ctx.Set("token", tokenString)
	}
}
