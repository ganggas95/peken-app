package middleware

import (
	"net/http"
	"peken-be/helper"
	"peken-be/models/web"
	"peken-be/repository"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if request has Authorization header
		authorizationHeader := ctx.GetHeader("Authorization")
		if authorizationHeader == "" {
			// Handle case when Authorization header is missing
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.Response(http.StatusUnauthorized, "Unauthorized", web.Null()))
			return
		}
		// Handle case when Authorization header is present
		// Split Bearer and token inside authorizationHeader
		if authorizationHeader[:7] != "Bearer " {
			// Handle case when Bearer is not present
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.Response(http.StatusUnauthorized, "Invalid Token Format. Use Bearer <token>", web.Null()))
			return
		}
		// Handle case when Bearer is present
		token := authorizationHeader[7:]
		user, err := helper.DecodeToken(token)
		if err != nil {
			// Handle case when token is invalid
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.Response(http.StatusUnauthorized, err.Error(), web.Null()))
			return
		}
		userInstance, err := userRepository.FindByID(uint(user["userId"].(float64)))
		if err != nil {
			// Handle case when user is not found
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.Response(http.StatusUnauthorized, "Invalid token", web.Null()))
			return
		}
		// Continue with the request
		// Set currentUser to the context
		ctx.Set("currentUser", userInstance)
		ctx.Copy().Next()
	}
}
