package test

import (
	"os"
	"peken-be/app"
	"peken-be/constants"
	"peken-be/controller"
	"peken-be/helper"
	"peken-be/models/domain"
	"peken-be/repository"
	"peken-be/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	memberRepository := repository.NewMemberRepository(db)
	passwordUtils := helper.NewPasswordUtils()
	validate := validator.New()

	userService := service.NewUserService(roleRepository, userRepository, passwordUtils, validate)
	userController := controller.NewUserController(userService)

	loginService := service.NewLoginService(userRepository, passwordUtils, validate)
	loginController := controller.NewLoginController(loginService)

	memberService := service.NewMemberService(memberRepository, validate)
	memberController := controller.NewMemberController(memberService)
	router := app.InitRoute(userController, loginController, userRepository, memberController)
	return router
}

func Init() {
	// app.InitLog()
	godotenv.Load("../.env")
	os.Setenv("ENV", "test")
	gin.SetMode(gin.ReleaseMode)
}

func InitRoleMock(db *gorm.DB) (domain.Role, domain.Role, domain.Role) {
	adminRole := domain.Role{
		Name: constants.ADMIN,
	}
	memberRole := domain.Role{
		Name: constants.MEMBER,
	}
	kasirRole := domain.Role{
		Name: constants.KASIR,
	}
	db.Create(&adminRole)
	db.Create(&memberRole)
	db.Create(&kasirRole)
	return adminRole, memberRole, kasirRole
}

func InitializeTestApp() (*gin.Engine, string) {
	Init()
	// Initialize db and router
	db := app.ConnectToDb()
	db.Logger.LogMode(logger.Info)
	router := InitRouter(db)

	// Clean up database and make migration
	db.Migrator().DropTable(&domain.User{})
	db.Migrator().DropTable(&domain.Role{})
	db.Migrator().DropTable("user_roles")
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Role{})
	// Initializer User data
	passwordUtils := helper.NewPasswordUtils()
	password, _ := passwordUtils.HashPassword("password")
	admin, _, _ := InitRoleMock(db)
	var roles []domain.Role
	roles = append(roles, admin)
	mockUser := domain.User{
		Username: "testuser",
		Password: password,
		Email:    "test@test.com",
		Roles:    roles,
	}
	mockUser2 := domain.User{
		Username: "testuser12",
		Password: password,
		Email:    "test12@test.com",
		Roles:    roles,
	}
	db.Create(&mockUser)
	db.Save(&mockUser)
	db.Create(&mockUser2)
	db.Save(&mockUser2)
	token, _ := helper.GenerateToken(&mockUser)
	return router, token
}
