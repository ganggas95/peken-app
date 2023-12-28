package service

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindAllUserRoles(ctx *gin.Context)
}
