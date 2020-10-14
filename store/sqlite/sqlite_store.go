package sqlite

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fareez-ahamed/todoapi/store"
)

// Store is the implmentation of Store using
// the Sqlite database
type Store struct {
	db *gorm.DB
}

// // GetTodos returns the list of todos from the database
// func (s *SqliteStore) GetTodos() ([]store.Todo, error) {

// }

// AddTodo adds a new Todo to the database
func (s *Store) AddTodo(todo *store.Todo) (*store.Todo, error) {
	result := s.db.Create(todo)
	if result.Error != nil {
		return nil, fmt.Errorf("adding a todo: %v", result.Error)
	}
	return todo, nil
}

// // MarkDone marks a todo as done or undone given its id
// func (s *SqliteStore) MarkDone(id uint, done bool) (store.Todo, error) {

// }

// // DeleteTodo soft deletes a Todo in the database
// func (s *SqliteStore) DeleteTodo(id uint) error {

// }

// NewStore returns a Store
func NewStore(file string) *Store {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&store.Todo{})

	return &Store{
		db: db,
	}
}
