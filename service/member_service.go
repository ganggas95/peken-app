package service

import "github.com/gin-gonic/gin"

type MemberService interface {
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
