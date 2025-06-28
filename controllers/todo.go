package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lef237/gin-mvs/models"
	"github.com/lef237/gin-mvs/services"
)

var todos []models.Todo
var idCounter = 1

func GetTodos(c *gin.Context) {
	if todos == nil {
		c.JSON(http.StatusOK, []models.Todo{})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, msg := newTodo.Validate(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	newTodo.ID = idCounter
	idCounter++
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func ToggleTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Toggle()
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func GetCompletedTodos(c *gin.Context) {
	completed := models.FilterCompleted(todos)
	c.JSON(http.StatusOK, completed)
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
