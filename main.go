// main.go
package main

import (
	"./auth"
	"./db"
	"./middleware"
	"./todo"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 连接数据库
	db.InitDB()
	defer db.CloseDB()

	// 设置HTML模板目录
	router.LoadHTMLGlob("templates/*")

	// 静态文件
	router.Static("/static", "./static")

	// 初始化路由
	initRoutes(router)

	// 启动服务器
	router.Run(":8080")
}

func initRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// 登录和注册
		api.POST("/login", auth.Login)
		api.POST("/register", auth.Register)

		// 使用认证中间件
		api.Use(middleware.AuthMiddleware())
		{
			// 待办事项相关操作
			api.GET("/todos", todo.GetTodos)
			api.POST("/todos", todo.AddTodo)
			api.PUT("/todos/:id", todo.UpdateTodo)
			api.DELETE("/todos/:id", todo.DeleteTodo)

			// 注销
			api.POST("/logout", auth.Logout)
		}
	}
}
