package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"log"

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
	case path == "/todo" && method == http.MethodGet:
		writeResponse(res, h.handleGetTodo(req))
	case path == "/todo" && method == http.MethodPost:
		writeResponse(res, h.handlePostTodo(req))
	default:
		writeResponse(res, &Response{
			Status: http.StatusNotFound,
			Error:  "Not Found",
		})
	}
}

func (h *TodoHandler) handleGetTodo(req *http.Request) *Response {
	var todos []Todo
	storeTodos, err := h.Store.GetTodos()
	if err != nil {
		return &Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
	}
	for _, t := range storeTodos {
		todos = append(todos, *mapTodo(&t))
	}
	return &Response{
		Status: http.StatusOK,
		Data:   todos,
	}
}

func (h *TodoHandler) handlePostTodo(req *http.Request) *Response {

	var data PostTodo

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		return &Response{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		}
	}

	todo, err := h.Store.AddTodo(&store.Todo{
		Description: data.Description,
	})
	if err != nil {
		return &Response{
			Error:  err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	return &Response{
		Status: http.StatusOK,
		Data:   mapTodo(todo),
	}
}

func stripEndSlash(path string) string {
	if strings.HasSuffix(path, "/") {
		return path[0 : len(path)-1]
	}
	return path
}

func writeResponse(writer http.ResponseWriter, response *Response) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(response.Status)
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("writing response: %v", err)
	}
	writer.Write(data)

}
