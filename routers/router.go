package routers

import (
	"matemask_login/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiAllRouterInit(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//前端POST路由组
	v2 := r.Group("/v2")
	{

		// 处理登录表单提交
		v2.POST("/login", controllers.LoginController{}.LoginHandler)

		// 显示个人资料页面
		v2.GET("/profile", controllers.ProfileController{}.ProfileHandler)
	}

}
