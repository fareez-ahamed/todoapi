package store

import (
	"database/sql"

	"gorm.io/gorm"
)

// Store interface represents store options
type Store interface {
	// GetTodos() ([]Todo, error)

	AddTodo(todo *Todo) (*Todo, error)
	// MarkDone(uint, bool) (Todo, error)
	// DeleteTodo(uint) error
}

// Todo model for storing a todo
type Todo struct {
	gorm.Model
	Description string
	Done        bool
	CompletedAt sql.NullTime
}
