package controller

import "github.com/gin-gonic/gin"

type LoginController interface {
	LoginAPI(ctx *gin.Context)
}
