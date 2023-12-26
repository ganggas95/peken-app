package app

import (
	"peken-be/controller"
	"peken-be/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(
	userController controller.UserController,
	loginController controller.LoginController,
) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.GlobalErrorHandler())

	routerGroup := router.Group("/api")
	// Login Routes
	routerGroup.POST("/login", loginController.LoginAPI)
	// User Routes
	user := routerGroup.Group("/users")
	{
		user.POST("", userController.Create)
		user.GET("", userController.FindAll)
		user.GET("/:userId", userController.FindById)
		user.PUT("/:userId", userController.Update)
		user.DELETE("/:userId", userController.Delete)
	}

	return router
}
