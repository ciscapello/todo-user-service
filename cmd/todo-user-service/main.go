package main

import (
	"fmt"
	"todo-user-service/internal/app/database"
	_ "todo-user-service/internal/app/database/migrations"
	"todo-user-service/internal/app/http"
	"todo-user-service/internal/config"
	"todo-user-service/internal/services"
)

func main() {
	fmt.Println("Application is runnin")
	// TODO: Get config
	conf := config.New()

	// TODO: Run Database
	pool := database.Connect()

	userService := services.NewUserService(pool)

	// TODO: Run HTTP server
	http.RunServer(conf.Port, userService)

	// TODO: Run gRPC server

	// TODO: Handle graceful shutdown
}
