package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(ctx *gin.Context, err string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  err,
	})
}
