package controller

import (
	"net/http"
	"peken-be/models/web"
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
	// defer helper.PanicHandler(ctx)
	loginResponse := ctrl.LoginService.Login(ctx)
	if loginResponse != nil {
		response := web.Response(http.StatusOK, "Success", loginResponse)
		ctx.JSON(http.StatusOK, response)
	}
}
