//go:build wireinject
// +build wireinject

package main

import (
	"peken-be/app"
	"peken-be/controller"
	"peken-be/helper"
	"peken-be/repository"
	"peken-be/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var loginSet = wire.NewSet(
	service.NewLoginService,
	wire.Bind(new(service.LoginService), new(*service.LoginServiceImpl)),
	controller.NewLoginController,
	wire.Bind(new(controller.LoginController), new(*controller.LoginControllerImpl)),
)
var passGenSet = wire.NewSet(
	helper.NewPasswordUtils,
	wire.Bind(new(helper.PasswordUtils), new(*helper.PasswordUtilsImpl)),
)

func InitializedServer() *gin.Engine {
	wire.Build(app.ConnectToDb, validator.New, userSet, passGenSet, loginSet, app.InitRoute)
	return nil
}
