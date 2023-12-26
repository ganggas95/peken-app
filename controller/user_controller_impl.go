package controller

import (
	"net/http"
	"peken-be/models/web"
	"peken-be/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

// NewUserController returns new UserController.
func NewUserController(userService service.UserService) *UserControllerImpl {

	return &UserControllerImpl{
		UserService: userService,
	}
}

func (ctrl *UserControllerImpl) Create(ctx *gin.Context) {
	// defer helper.PanicHandler(ctx)
	user := ctrl.UserService.Save(ctx)
	if user != nil {
		response := web.Response(http.StatusCreated, "Success", user)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (ctrl *UserControllerImpl) Update(ctx *gin.Context) {
	// defer helper.PanicHandler(ctx)
	user := ctrl.UserService.Update(ctx)
	if user != nil {
		response := web.Response(http.StatusOK, "Success", user)
		ctx.JSON(http.StatusOK, response)
	}
}

func (ctrl *UserControllerImpl) Delete(ctx *gin.Context) {
	// defer helper.PanicHandler(ctx)
	userId := ctrl.UserService.Delete(ctx)
	if userId != nil {
		response := web.Response(http.StatusOK, "Success", web.Null())
		ctx.JSON(http.StatusOK, response)
	}
}

func (ctrl *UserControllerImpl) FindById(ctx *gin.Context) {
	// defer helper.PanicHandler(ctx)
	ctrl.UserService.FindByID(ctx)
}

func (ctrl *UserControllerImpl) FindAll(ctx *gin.Context) {
	// defer helper.PanicHandler(ctx)
	users := ctrl.UserService.FindAll(ctx)
	response := web.Response(http.StatusOK, "Success", users)

	ctx.JSON(http.StatusOK, response)
}
