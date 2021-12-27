package main

import (
	"back_server/db"
	"back_server/handle"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//res := handle.EncodePassword("admin")
	//fmt.Println(res)
	//return
	// 初始化数据库
	err := db.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败", err)
		return
	}
	// 路由
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("pong"))
	})
	r.POST("/login", handle.Login)
	r.POST("/register", handle.Register)

	// 运行服务
	err = r.Run(":3748")
	if err != nil {
		fmt.Println("运行服务失败", err)
		return
	}
}
