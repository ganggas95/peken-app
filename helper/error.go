package helper

import (
	"fmt"
	"net/http"
	"peken-be/models/web"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := ConvertStringToInt(strArr[0])
		msg := strings.Trim(strArr[1], " ")
		switch key {
		case http.StatusBadRequest:
			c.JSON(http.StatusBadRequest, web.Response(key, msg, web.Null()))
			c.Abort()
		case http.StatusUnauthorized:
			c.JSON(http.StatusUnauthorized, web.Response(key, msg, web.Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, web.Response(key, msg, web.Null()))
			c.Abort()
		}
	}
}
