package models

import "testing"

func TestValidate(t *testing.T) {
	todo := Todo{Title: "", Completed: false}
	ok, msg := todo.Validate()
	if ok {
		t.Error("Validate should fail for empty title")
	}
	if msg == "" {
		t.Error("Validate should return error message for empty title")
	}

	longTitle := ""
	for i := 0; i < 101; i++ {
		longTitle += "a"
	}
	todo.Title = longTitle
	ok, msg = todo.Validate()
	if ok {
		t.Error("Validate should fail for title exceeding 100 characters")
	}
	if msg == "" {
		t.Error("Validate should return error message for long title")
	}

	todo.Title = "test"
	ok, _ = todo.Validate()
	if !ok {
		t.Error("Validate should pass for non-empty title")
	}
}

func TestToggle(t *testing.T) {
	todo := Todo{Title: "toggle", Completed: false}
	todo.Toggle()
	if !todo.Completed {
		t.Error("Toggle should set Completed to true")
	}
	todo.Toggle()
	if todo.Completed {
		t.Error("Toggle should set Completed to false")
	}
}

func TestFilterCompleted(t *testing.T) {
	todos := []Todo{
		{Title: "a", Completed: true},
		{Title: "b", Completed: false},
		{Title: "c", Completed: true},
	}
	completed := FilterCompleted(todos)
	if len(completed) != 2 {
		t.Errorf("expected 2 completed todos, got %d", len(completed))
	}
}
