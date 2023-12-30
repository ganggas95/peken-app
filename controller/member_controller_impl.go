package controller

import (
	"peken-be/service"

	"github.com/gin-gonic/gin"
)

type MemberControllerImpl struct {
	MemberService service.MemberService
}

func NewMemberController(
	memberService service.MemberService,
) *MemberControllerImpl {
	return &MemberControllerImpl{
		MemberService: memberService,
	}
}

func (controller *MemberControllerImpl) FindAll(ctx *gin.Context) {

}

func (controller *MemberControllerImpl) Create(ctx *gin.Context) {
}

func (controller *MemberControllerImpl) Update(ctx *gin.Context) {
}

func (controller *MemberControllerImpl) Delete(ctx *gin.Context) {
}

func (controller *MemberControllerImpl) FindById(ctx *gin.Context) {
}
