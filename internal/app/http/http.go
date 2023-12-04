package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	auth "todo-user-service/internal/app/http/handlers"
	"todo-user-service/internal/app/http/middleware"
)

func RunServer(port string) {
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	//r.Use(middleware.Logger)

	if err := auth.CreateAuthHandler(r); err != nil {
		log.Fatal("failed to creating auth handler")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
