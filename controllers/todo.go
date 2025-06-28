package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lef237/gin-mvs/models"
	"github.com/lef237/gin-mvs/services"
)

func GetTodos(c *gin.Context) {
	todos := services.GetTodos()
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, msg, ok := services.CreateTodo(newTodo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func ToggleTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, ok := services.ToggleTodo(id)
	if ok {
		c.JSON(http.StatusOK, todo)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func GetCompletedTodos(c *gin.Context) {
	completed := services.GetCompletedTodos()
	c.JSON(http.StatusOK, completed)
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if services.DeleteTodo(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
