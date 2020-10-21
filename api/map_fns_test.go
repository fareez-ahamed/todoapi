package api

import (
	"testing"

	"github.com/fareez-ahamed/todoapi/store"
)

func TestMapTodo(t *testing.T) {
	type testData struct {
		StoreTodo store.Todo
		MapTodo   Todo
	}
	data := []testData{
		{store.Todo{Description: "Test1", Done: false}, Todo{Description: "Test1", Done: false}},
		{store.Todo{Description: "Another", Done: true}, Todo{Description: "Another", Done: true}},
	}
	for _, rec := range data {
		mappedTodo := mapTodo(&rec.StoreTodo)
		if mappedTodo.Description != rec.MapTodo.Description ||
			mappedTodo.Done != rec.MapTodo.Done {
			t.Errorf("Expected: %v\nGot: %v", rec.MapTodo, mappedTodo)
		}
	}
}
