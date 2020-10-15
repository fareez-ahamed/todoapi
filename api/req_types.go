package api

// PostTodo is the request data for adding a
// new todo
type PostTodo struct {
	Description string `json:"description"`
}

// Response is the type of standard response object
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

// Todo type which will be returned in the response
type Todo struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
