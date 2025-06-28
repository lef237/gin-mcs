package services

import (
	"testing"

	"github.com/lef237/gin-mcs/models"
)

func TestCreateAndGetTodos(t *testing.T) {
	Todos = []models.Todo{}
	IdCounter = 1

	todo := models.Todo{Title: "test todo", Completed: false}
	created, msg, ok := CreateTodo(todo)
	if !ok {
		t.Fatalf("CreateTodo failed: %s", msg)
	}
	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}

	list := GetTodos()
	if len(list) != 1 {
		t.Errorf("expected 1 todo, got %d", len(list))
	}
	if list[0].Title != "test todo" {
		t.Errorf("expected title 'test todo', got '%s'", list[0].Title)
	}
}

func TestToggleTodo(t *testing.T) {
	Todos = []models.Todo{}
	IdCounter = 1
	created, _, _ := CreateTodo(models.Todo{Title: "toggle me"})
	_, ok := ToggleTodo(created.ID)
	if !ok {
		t.Fatal("ToggleTodo failed")
	}
	if !Todos[0].Completed {
		t.Error("expected todo to be completed after toggle")
	}
}

func TestDeleteTodo(t *testing.T) {
	Todos = []models.Todo{}
	IdCounter = 1
	created, _, _ := CreateTodo(models.Todo{Title: "delete me"})
	ok := DeleteTodo(created.ID)
	if !ok {
		t.Fatal("DeleteTodo failed")
	}
	if len(Todos) != 0 {
		t.Error("expected todos to be empty after delete")
	}
}

func TestDeleteNonExistentTodo(t *testing.T) {
	Todos = []models.Todo{}
	IdCounter = 1

	// idが1のtodoを作成
	created1, _, _ := CreateTodo(models.Todo{Title: "todo 1"})
	if created1.ID != 1 {
		t.Errorf("expected ID 1, got %d", created1.ID)
	}

	// idが2のtodoを作成
	created2, _, _ := CreateTodo(models.Todo{Title: "todo 2"})
	if created2.ID != 2 {
		t.Errorf("expected ID 2, got %d", created2.ID)
	}

	// idが1のtodoを削除
	ok := DeleteTodo(created1.ID)
	if !ok {
		t.Fatal("DeleteTodo failed for ID 1")
	}

	// 再度idが1のtodoを削除しようとすると失敗することを確認
	ok = DeleteTodo(created1.ID)
	if ok {
		t.Fatal("DeleteTodo should have failed for non-existent ID 1")
	}
}
