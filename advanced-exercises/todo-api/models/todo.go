package models

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func NewTodo() Todo {
	return Todo{}
}

func IsZero(todo Todo) bool {
	return todo == Todo{}
}

type UpdateTodoRequest struct {
	Text *string `json:"text"`
	Done *bool   `json:"done"`
}

func (t UpdateTodoRequest) CopyTo(tIn *Todo) {
	tIn.Done = *t.Done
	tIn.Text = *t.Text
}

type CreateTodoRequest struct {
	Text *string `json:"text"`
	Done *bool   `json:"done"`
}

func (t CreateTodoRequest) CopyTo(tIn *Todo) {
	tIn.Done = *t.Done
	tIn.Text = *t.Text
}
