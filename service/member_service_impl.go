package service

import (
	"fmt"
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
	members, _ := service.MemberRepository.FindAll()
	fmt.Println(members)
}
