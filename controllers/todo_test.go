package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lef237/gin-mcs/models"
	"github.com/lef237/gin-mcs/services"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos", GetTodos)
	r.POST("/todos", CreateTodo)
	r.PUT("/todos/:id/toggle", ToggleTodo)
	r.GET("/todos/completed", GetCompletedTodos)
	r.DELETE("/todos/:id", DeleteTodo)
	return r
}

func TestCreateAndGetTodos(t *testing.T) {
	services.Todos = []models.Todo{} // サービスの状態を初期化
	services.IdCounter = 1

	r := setupRouter()
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"title":"controller test"}`)
	req, _ := http.NewRequest("POST", "/todos", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/todos", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var todos []models.Todo
	if err := json.Unmarshal(w.Body.Bytes(), &todos); err != nil {
		t.Fatal("invalid response json")
	}
	if len(todos) != 1 || todos[0].Title != "controller test" {
		t.Error("unexpected todos response")
	}
}
