package http

import (
	"fmt"
	"log"
	"net/http"
	"todo-user-service/internal/app/http/handlers/auth"
	"todo-user-service/internal/app/http/middleware"
	"todo-user-service/internal/services"

	"github.com/gorilla/mux"
)

func RunServer(port string, userService *services.UserService) {
	r := mux.NewRouter()
	r.Use(middleware.Middleware)

	err := auth.CreateAuthHandler(r, userService)
	if err != nil {
		log.Fatal("failed to creating auth handler")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
