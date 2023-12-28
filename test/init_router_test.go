package test

import (
	"os"
	"peken-be/app"
	"peken-be/controller"
	"peken-be/helper"
	"peken-be/models/domain"
	"peken-be/repository"
	"peken-be/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	userRepository := repository.NewUserRepository(db)
	passwordUtils := helper.NewPasswordUtils()
	validate := validator.New()
	userService := service.NewUserService(userRepository, passwordUtils, validate)
	userController := controller.NewUserController(userService)
	loginService := service.NewLoginService(userRepository, passwordUtils, validate)
	loginController := controller.NewLoginController(loginService)
	router := app.InitRoute(userController, loginController, userRepository)
	return router
}

func Init() {
	// app.InitLog()
	godotenv.Load("../.env")
	os.Setenv("ENV", "test")
}

func InitializeTestApp() (*gin.Engine, *domain.User) {
	Init()
	// Initialize db and router
	db := app.ConnectToDb()
	router := InitRouter(db)

	// Clean up database and make migration
	db.Migrator().DropTable(&domain.User{})
	db.AutoMigrate(&domain.User{})
	// Initializer User data
	passwordUtils := helper.NewPasswordUtils()
	password, _ := passwordUtils.HashPassword("password")
	mockUser := domain.User{
		Username: "testuser",
		Password: password,
	}
	db.Create(&mockUser)
	return router, &mockUser
}
