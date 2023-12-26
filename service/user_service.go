package service

import (
	"peken-be/models/domain"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Save(ctx *gin.Context) *domain.User
	Update(ctx *gin.Context) *domain.User
	Delete(ctx *gin.Context) *int
	FindByID(ctx *gin.Context) *domain.User
	FindAll(ctx *gin.Context) []domain.User
}
