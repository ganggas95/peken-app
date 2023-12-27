package test

import (
	"peken-be/app"
	"peken-be/controller"
	"peken-be/helper"
	"peken-be/repository"
	"peken-be/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	userRepositoryImpl := repository.NewUserRepository(db)
	passwordUtilsImpl := helper.NewPasswordUtils()
	validate := validator.New()
	userService := service.NewUserService(userRepositoryImpl, passwordUtilsImpl, validate)
	userController := controller.NewUserController(userService)
	loginServiceImpl := service.NewLoginService(userRepositoryImpl, passwordUtilsImpl, validate)
	loginController := controller.NewLoginController(loginServiceImpl)
	router := app.InitRoute(userController, loginController)
	return router
}
