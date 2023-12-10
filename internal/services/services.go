package services

import "github.com/jackc/pgx/v5/pgxpool"

type Services struct {
	UserService  *UserService
	TodosService *TodosService
}

func Init(pool *pgxpool.Pool) *Services {
	return &Services{
		UserService:  NewUserService(pool),
		TodosService: NewTodosService(pool),
	}
}
