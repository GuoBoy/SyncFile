package handle

import (
	"back_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context)  {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ErrorResponse(ctx, "参数不足")
		return
	}
	if err := ExistUser(&user); err!=nil{
		//fmt.Println("查询用户是否存在", err)
		//fmt.Println("用户不存在，添加")
		user.Token = NewToken(user.Username)
		if err = AddUser(&user); err!=nil{
			ErrorResponse(ctx, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "register ok",
			"token": user.Token,
		})
		return
	}
	ErrorResponse(ctx, "用户已存在")
}
