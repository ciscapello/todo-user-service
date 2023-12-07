package auth

import (
	"net/http"
	"todo-user-service/internal/services"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	router      *mux.Router
	userService *services.UserService
}

func CreateAuthHandler(router *mux.Router, userService *services.UserService) error {
	handler := AuthHandler{router: router, userService: userService}

	handler.router.HandleFunc("/auth/register", handler.Registration).Methods(http.MethodPost)
	handler.router.HandleFunc("/auth/login", handler.Login).Methods(http.MethodPost)
	return nil
}
