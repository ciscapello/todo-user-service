package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectsRepo struct {
	pool *pgxpool.Pool
}

func NewProjectRepo(pool *pgxpool.Pool) *ProjectsRepo {
	return &ProjectsRepo{
		pool: pool,
	}
}

func (r *UsersRepo) CreateProject(email string, password string) (int, string) {
	sqlStatement := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	RETURNING id
	`

	var id int
	err := r.pool.QueryRow(context.Background(), sqlStatement, email, password).Scan(&id)
	if err != nil {
		fmt.Printf("error %v", err)
		return 0, "cannot create user in db"
	}
	return id, ""
}
