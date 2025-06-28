package services

import (
	"github.com/lef237/gin-mcs/models"
)

var Todos []models.Todo = []models.Todo{}
var IdCounter = 1

func GetTodos() []models.Todo {
	if Todos == nil {
		return []models.Todo{}
	}
	return Todos
}

func CreateTodo(newTodo models.Todo) (models.Todo, string, bool) {
	if ok, msg := newTodo.Validate(); !ok {
		return models.Todo{}, msg, false
	}
	newTodo.ID = IdCounter
	IdCounter++
	Todos = append(Todos, newTodo)
	return newTodo, "", true
}

func ToggleTodo(id int) (models.Todo, bool) {
	for i := range Todos {
		if Todos[i].ID == id {
			Todos[i].Toggle()
			return Todos[i], true
		}
	}
	return models.Todo{}, false
}

func GetCompletedTodos() []models.Todo {
	return models.FilterCompleted(Todos)
}

func DeleteTodo(id int) bool {
	for i, t := range Todos {
		if t.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return true
		}
	}
	return false
}
