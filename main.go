package main

import (
	"net/http"

	"github.com/fareez-ahamed/todoapi/api"
	"github.com/fareez-ahamed/todoapi/store/sqlite"
)

func main() {
	store := sqlite.NewStore("test.db")
	todoHandler := &api.TodoHandler{
		Store: store,
	}
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: todoHandler,
	}
	httpServer.ListenAndServe()
}
