package main

import (
	"net/http"
	"strings"

	"github.com/fareez-ahamed/todoapi/store"
)

// TodoHandler represents the Todo api requests handling
type TodoHandler struct {
	Store store.Store
}

func (h *TodoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := stripEndSlash(req.URL.Path)
	method := req.Method
	switch {
	case path == "/todo" && method == "GET":
		h.handleGetTodo(res, req)
	default:
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Not Found!"))
	}
}

func (h *TodoHandler) handleGetTodo(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Alhamdhulillah"))
}

func (h *TodoHandler) handlePostTodo(res http.ResponseWriter, req *http.Request) {

}

func stripEndSlash(path string) string {
	if strings.HasSuffix(path, "/") {
		return path[0 : len(path)-1]
	}
	return path
}
