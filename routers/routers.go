package routers

import (
	"bubble/controller"
	"bubble/setting"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	//github.com/War11/Memorandum_GinandGorm.git
	r.LoadHTMLGlob("/go/src/github.com/War11/Memorandum_GinandGorm/templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{

		v1Group.POST("/todo", controller.CreateTodo)

		v1Group.GET("/todo", controller.GetTodoList)

		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
