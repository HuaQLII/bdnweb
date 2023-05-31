package controllers

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserController 将抽离的出来的接口返回值
// 放到这里并挂载到UserController结构体上,实现接口的继承
// 那么在接口那边方法引用时要先实例化这个结构体，然后再调用方法
// （查看的话请跳转到userRouter看具体使用）
type ProfileController struct {
	// 这里可以放一些公共的方法比如自定状态码等
}

// 用户登录
func (PC ProfileController)ProfileHandler(c *gin.Context) {
	// 获取会话中的地址
	session := sessions.Default(c)
	address := session.Get("address")

	// 如果地址不存在，则重定向到登录页面
	if address == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// 显示个人资料页面
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"address": address,
	})
}