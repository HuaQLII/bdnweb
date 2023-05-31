package main

import (
	"matemask_login/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//路由初始化
	router := gin.Default()

	// 设置会话存储
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.LoadHTMLGlob("templates/*")
	routers.ApiAllRouterInit(router)

	// 显示登录页面

	// 启动服务器
	router.Run(":8080")
}
