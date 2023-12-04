package main

import (
	"fmt"
	"todo-user-service/internal/app/http"
	"todo-user-service/internal/config"
)

func main() {
	fmt.Println("Application is runnin")
	// TODO: Get config
	conf := config.New()

	// TODO: Run Database

	// TODO: Run HTTP server
	http.RunServer(conf.Port)

	// TODO: Run gRPC server

	// TODO: Handle graceful shutdown
}
