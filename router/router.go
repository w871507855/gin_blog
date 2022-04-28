package router

import (
	"gin_learning/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)
	e.GET("/", controller.Index)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	// 博客列表
	e.GET("/post_index", controller.GetPostIndex)
	// 添加博客
	e.POST("/post", controller.AddPost)
	// 跳转到添加博客页面
	e.GET("/post", controller.GoAddPost)

	// 跳转到博客详情页面
	e.GET("/detail", controller.PostDetail)

	e.Run()
}
