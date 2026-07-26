package router

import (
	"memo/controller"
	"memo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 注册路由
	r.POST("/api/register", controller.Register)

	r.POST("/api/login", controller.Login)

	authrouter := r.Group("/api")
	authrouter.Use(middleware.AuthMiddleware())
	{
		authrouter.POST("/todos", controller.CreateTodo)
		authrouter.GET("/todos", controller.GetTodoList)
		authrouter.PUT("/todos", controller.UpdateTodo)
		authrouter.DELETE("/todos", controller.DeleteTodo)
	}
	return r

}
