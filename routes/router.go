package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lef237/gin-mvc/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	r.GET("/todos/completed", controllers.GetCompletedTodos)
	r.PUT("/todos/:id/toggle", controllers.ToggleTodo)

	return r
}
