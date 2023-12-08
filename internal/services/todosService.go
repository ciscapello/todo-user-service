package services

import "github.com/jackc/pgx/v5/pgxpool"

type TodosService struct {
	pool *pgxpool.Pool
}

func NewTodosService(dbPool *pgxpool.Pool) *TodosService {
	return &TodosService{pool: dbPool}
}

func (s *TodosService) CreateTodo() {}
