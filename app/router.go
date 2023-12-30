package app

import (
	"os"
	"peken-be/controller"
	"peken-be/middleware"
	"peken-be/repository"

	"github.com/gin-gonic/gin"
)

func InitRoute(
	userController controller.UserController,
	loginController controller.LoginController,
	userRepository repository.UserRepository,
	memberController controller.MemberController,
) *gin.Engine {
	router := gin.New()
	if os.Getenv("ENV") != "test" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	}
	router.Use(middleware.GlobalErrorHandler())

	routerGroup := router.Group("/api")
	// Login Routes
	routerGroup.POST("/login", loginController.LoginAPI)
	// User Routes
	userRoute := routerGroup.Group("/users")
	userRoute.Use(middleware.AuthMiddleware(userRepository))
	userRoute.POST("", userController.Create)
	userRoute.GET("", userController.FindAll)
	userRoute.GET("/:userId", userController.FindById)
	userRoute.PUT("/:userId", userController.Update)
	userRoute.DELETE("/:userId", userController.Delete)
	userRoute.GET("/roles", userController.FindAllUserRoles)

	memberRoutes := routerGroup.Group("/members")
	memberRoutes.POST("", memberController.Create)
	memberRoutes.GET("", memberController.FindAll)
	memberRoutes.GET("/:memberId", memberController.FindById)
	memberRoutes.PUT("/:memberId", memberController.Update)
	memberRoutes.DELETE("/:memberId", memberController.Delete)
	return router
}
