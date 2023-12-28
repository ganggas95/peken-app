package middleware

import (
	"net/http"
	errors "peken-be/models/errors"
	"peken-be/models/web"

	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.ByType(gin.ErrorTypeAny).Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *errors.LudesError:
				c.AbortWithStatusJSON(e.Status, web.Response(e.Status, e.Message, web.Null()))
			default:
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					web.Response(http.StatusInternalServerError, err.Error(), web.Null()))
			}
		}
	}
}
