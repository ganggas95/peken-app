package controller

import (
	"fmt"
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

	ctrl.UserService.Save(ctx)
}

func (ctrl *UserControllerImpl) Update(ctx *gin.Context) {

	ctrl.UserService.Update(ctx)

}

func (ctrl *UserControllerImpl) Delete(ctx *gin.Context) {

	ctrl.UserService.Delete(ctx)

}

func (ctrl *UserControllerImpl) FindById(ctx *gin.Context) {

	ctrl.UserService.FindByID(ctx)
}

func (ctrl *UserControllerImpl) FindAll(ctx *gin.Context) {
	fmt.Println(ctx.Get("currentUser"))
	ctrl.UserService.FindAll(ctx)
}
