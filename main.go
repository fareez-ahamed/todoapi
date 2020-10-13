package main

import (
	"net/http"

	"github.com/fareez-ahamed/todoapi/store/sqlite"
)

// APIServer represents the Rest API Server
type APIServer struct {
	Port    int
	Handler http.Handler
}

func main() {
	store := sqlite.NewStore("test.db")
	todoHandler := TodoHandler{
		Store: store,
	}
	api := APIServer{
		Port:    8080,
		Handler: &todoHandler,
	}
	api.Listen()
}

// Listen starts the api server
func (s *APIServer) Listen() {
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: s.Handler,
	}
	httpServer.ListenAndServe()
}
