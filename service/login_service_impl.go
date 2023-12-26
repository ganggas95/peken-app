package service

import (
	"net/http"
	"peken-be/helper"
	"peken-be/models/errors"
	"peken-be/models/web"
	"peken-be/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type LoginServiceImpl struct {
	PasswordUtils  helper.PasswordUtils
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// NewLoginService returns new LoginService.
func NewLoginService(
	userRepository repository.UserRepository,
	passwordUtils helper.PasswordUtils,
	validate *validator.Validate,
) *LoginServiceImpl {
	return &LoginServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
		PasswordUtils:  passwordUtils,
	}
}

func (loginService *LoginServiceImpl) Login(ctx *gin.Context) {
	var request web.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		cError := errors.NewLudesError(http.StatusBadRequest, err.Error())
		ctx.Error(cError)
		return
	}
	if err := loginService.Validate.Struct(request); err != nil {
		cError := errors.NewLudesError(http.StatusBadRequest, err.Error())
		ctx.Error(cError)
		return
	}
	user, err := loginService.UserRepository.FindByUsername(request.Username)
	if err != nil {
		cError := errors.NewLudesError(http.StatusUnauthorized, "Username dan password salah")
		ctx.Error(cError)
		return
	}
	passwordMatch := loginService.PasswordUtils.CheckPasswordHash(request.Password, user.Password)
	if !passwordMatch {
		cError := errors.NewLudesError(http.StatusUnauthorized, "Username dan password salah")
		ctx.Error(cError)
		return
	} else {
		loginResponse := web.NewLoginResponse("", "")
		response := web.Response(http.StatusOK, "Success", loginResponse)
		ctx.JSON(http.StatusOK, response)
	}
}
