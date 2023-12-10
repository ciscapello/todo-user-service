package main

import (
	"fmt"
	"todo-user-service/internal/app"
	_ "todo-user-service/internal/app/database/migrations"
)

func main() {
	fmt.Println("Application is runnin")
	// TODO: Get config
	app := app.Init()
	app.Run()

	// conf := config.New()

	// // TODO: Run Database
	// pool := database.Connect()

	// userService := services.NewUserService(pool)
	// authService := services.NewAuthService()

	// // // TODO: Run HTTP server
	// http.RunServer(conf.Port, userService, authService)

	// TODO: Run gRPC server

	// TODO: Handle graceful shutdown
}
