package controller

import (
	"peken-be/service"

	"github.com/gin-gonic/gin"
)

type LoginControllerImpl struct {
	LoginService service.LoginService
}

// NewUserController returns new LoginController.
func NewLoginController(loginService service.LoginService) *LoginControllerImpl {

	return &LoginControllerImpl{
		LoginService: loginService,
	}
}

func (ctrl *LoginControllerImpl) LoginAPI(ctx *gin.Context) {
	ctrl.LoginService.Login(ctx)
}
