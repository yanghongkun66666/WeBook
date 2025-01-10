package main

import (
	"WEBBOOK/webook/internal/web"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//u := &web.UserHandler{}
	u := web.NewUserHandler()
	u.RegisterRoutesV1(server.Group("/users"))
	//这样的好处是 我能看到我的前缀已经被占用了，业务不会冲突

	//u.RegisterRoutes(server)
	server.Run(":8080")
	//
	//以上是使用init_web的方式
}
