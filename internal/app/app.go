package app

import (
	"database/sql"
	"net/http"
	"todo-user-service/internal/app/database"
	"todo-user-service/internal/app/httpServer"
	"todo-user-service/internal/config"
	"todo-user-service/internal/services"
	"todo-user-service/pkg/tokenManager"
)

type App struct {
	httpServer http.Server
	database   sql.DB
}

func Init() *App {
	return &App{}
}

func (app *App) Run() {
	conf := config.New()

	// TODO: Run Database
	pool := database.Connect()

	services := services.Init(pool)

	tokenManager := tokenManager.NewTokenManager()

	// userService := services.NewUserService(pool)
	// authService := services.NewAuthService()

	// // TODO: Run HTTP server

	httpServer.RunServer(conf.Port, services, tokenManager)

	// httpServer.RunServer(conf.Port, services, tokenManager)

}
