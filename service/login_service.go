package service

import (
	"peken-be/models/web"

	"github.com/gin-gonic/gin"
)

type LoginService interface {
	// Login
	Login(ctx *gin.Context) *web.LoginResponse
}
