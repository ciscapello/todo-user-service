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

func (s *UserService) CreateUser(username string, password string) (int, string) {
	isUserExists := s.checkUserIsExists(username)
	if isUserExists {
		return 0, "user is already exists"
	}

	sqlStatement := `
	INSERT INTO users (username, password)
	VALUES ($1, $2)
	RETURNING id`

	var id int
	err := s.pool.QueryRow(context.Background(), sqlStatement, username, password).Scan(&id)
	if err != nil {
		fmt.Printf("error %v", err)
		return 0, "cannot create user in db"
	}

	return id, ""
}

func (s *UserService) checkUserIsExists(username string) bool {
	sqlStatement := `
	SELECT id FROM users WHERE username=$1 LIMIT 1
	`
	var id int
	err := s.pool.QueryRow(context.Background(), sqlStatement, username).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
	return id != 0
}
