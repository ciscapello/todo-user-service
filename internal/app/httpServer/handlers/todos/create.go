package todos

import (
	"fmt"
	"net/http"
)

type createTodoRequestBody struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Priority string `json:"priority"`
	Deadline string `json:"string"`
}

func (h *TodosHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var body createTodoRequestBody
	fmt.Println(body)
}
