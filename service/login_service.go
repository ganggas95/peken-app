package service

import (
	"github.com/gin-gonic/gin"
)

type LoginService interface {
	// Login
	Login(ctx *gin.Context)
}
