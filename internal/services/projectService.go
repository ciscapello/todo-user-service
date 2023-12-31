package services

import (
	"todo-user-service/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectService struct {
	projectsRepo *repository.ProjectsRepo
	pool         *pgxpool.Pool
}

func NewProjectService(projectRepo *repository.ProjectsRepo, dbPool *pgxpool.Pool) *ProjectService {
	return &ProjectService{projectsRepo: projectRepo, pool: dbPool}
}

func (s *ProjectService) CreateProject() {}
