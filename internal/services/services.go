package services

import (
	"todo-user-service/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Services struct {
	UserService  *UserService
	TodosService *TodosService
}

func Init(pool *pgxpool.Pool) *Services {
	usersRepo := repository.NewUserRepo(pool)
	sessionsRepo := repository.NewSessionsRepo(pool)
	return &Services{
		UserService:  NewUserService(usersRepo, sessionsRepo),
		TodosService: NewTodosService(pool),
	}
}
