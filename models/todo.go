package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Todo) Toggle() {
	t.Completed = !t.Completed
}

func (t *Todo) Validate() (bool, string) {
	if len(t.Title) == 0 {
		return false, "Title is required"
	}
	if len(t.Title) > 100 {
		return false, "Title must be 100 characters or less"
	}
	return true, ""
}

func FilterCompleted(todos []Todo) []Todo {
	var result []Todo
	for _, t := range todos {
		if t.Completed {
			result = append(result, t)
		}
	}
	return result
}
