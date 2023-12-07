package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	pool *pgxpool.Pool
}

func NewUserService(dbPool *pgxpool.Pool) *UserService {
	return &UserService{pool: dbPool}
}

func (s *UserService) CreateUser(phone string, password string) (int, error) {
	sqlStatement := `
	INSERT INTO users (phone, password)
	VALUES ($1, $2)
	RETURNING id`
	var id int
	err := s.pool.QueryRow(context.Background(), sqlStatement, phone, password).Scan(&id)
	if err != nil {
		fmt.Printf("error %v", err)
		return 0, err
	}

	return id, nil
}
