package auth

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	router *mux.Router
}

func CreateAuthHandler(router *mux.Router) error {
	handler := Handler{router: router}

	handler.router.HandleFunc("/auth/register", handler.Registration).Methods(http.MethodPost)
	return nil
}
