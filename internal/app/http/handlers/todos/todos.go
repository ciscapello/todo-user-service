package todos

import (
	"net/http"
	"todo-user-service/internal/services"

	"github.com/gorilla/mux"
)

type TodosHandler struct {
	router       *mux.Router
	todosService *services.TodosService
}

func CreateTodosHandler(router *mux.Router, todosService *services.TodosService) error {
	handler := TodosHandler{router: router, todosService: todosService}

	handler.router.HandleFunc("/todos", handler.Create).Methods(http.MethodPost)
	return nil
}
