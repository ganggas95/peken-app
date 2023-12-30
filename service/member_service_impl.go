package service

import (
	"peken-be/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type MemberServiceImpl struct {
	MemberRepository repository.MemberRepository
	Validate         *validator.Validate
}

// NewMemberService returns new MemberService.
func NewMemberService(
	memberRepository repository.MemberRepository,
	validate *validator.Validate,
) *MemberServiceImpl {

	return &MemberServiceImpl{
		MemberRepository: memberRepository,
		Validate:         validate,
	}
}

func (service *MemberServiceImpl) FindAll(ctx *gin.Context) {
}

func (service *MemberServiceImpl) Save(ctx *gin.Context) {
}

func (service *MemberServiceImpl) Update(ctx *gin.Context) {
}

func (service *MemberServiceImpl) Delete(ctx *gin.Context) {
}

func (service *MemberServiceImpl) FindByID(ctx *gin.Context) {
}
