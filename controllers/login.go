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
type LoginController struct {
	// 这里可以放一些公共的方法比如自定状态码等
}

// 用户登录
func (LC LoginController) LoginHandler(c *gin.Context) {
	// 获取表单数据
	address := c.PostForm("address")

	// 验证地址是否合法
	// 在这里，您可以使用以太坊的web3.js库来验证地址是否有效。

	// 如果地址有效，则将其保存到会话中
	session := sessions.Default(c)
	session.Set("address", address)
	session.Save()

	// 重定向到用户的个人资料页面
	c.Redirect(http.StatusSeeOther, "/v2/profile")
}
