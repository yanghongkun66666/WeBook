package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	registerUsersRoutes(server)

	return server
}

// 实际上这个函数也会非常大
// 有些人会进一步拆分
func registerUsersRoutes(server *gin.Engine) {
	//进一步拆分
	u := UserHandler{}
	//u：这是一个新的变量，它的类型是 *web.UserHandler（即指向 web.UserHandler 结构体的指针）

	//注册
	server.POST("/users/signup", u.SignUp)

	//登录
	server.POST("/users/login", u.Login)

	//编辑用户
	server.POST("/users/edit", u.Edit)

	server.GET("users/profile", u.Profile)
}
