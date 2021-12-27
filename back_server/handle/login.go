package handle

import (
	"back_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context)  {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}
	if err := ValidUser(&user); err!=nil{
		ErrorResponse(ctx, err.Error())
		return
	}
	token := NewToken(user.Username)
	if err := UpdateToken(&user, token); err!=nil{
		ErrorResponse(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "login ok",
		"token": token,
	})
}
