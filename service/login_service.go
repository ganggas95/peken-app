package service

import (
	"peken-be/models/domain"

	"github.com/gin-gonic/gin"
)

type LoginService interface {
	// Login
	Login(ctx *gin.Context)
	GenerateToken(user *domain.User) (string, error)
	DecodeToken(token string) (*domain.User, error)
}
