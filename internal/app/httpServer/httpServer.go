package httpServer

import (

	// "todo-user-service/internal/app/http/handlers/auth"

	"todo-user-service/internal/app/httpServer/handlers"
	"todo-user-service/internal/services"
	"todo-user-service/pkg/tokenManager"

	"github.com/gin-gonic/gin"
)

func RunServer(port string, services *services.Services, tokenManager *tokenManager.TokenManager) {
	// r := mux.NewRouter()
	r := gin.Default()

	handler := handlers.NewHandler(services, tokenManager, r)

	handler.RunHandler()

	r.Run()
	// r.Use(middleware.Middleware)

	// err := auth.CreateAuthHandler(r, userService)
	// if err != nil {
	// 	log.Fatal("failed to creating auth handler")
	// }

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
