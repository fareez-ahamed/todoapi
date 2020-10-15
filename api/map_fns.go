package api

import "github.com/fareez-ahamed/todoapi/store"

func mapTodo(todo *store.Todo) *Todo {
	return &Todo{
		ID:          todo.ID,
		Description: todo.Description,
		Done:        todo.Done,
	}
}
