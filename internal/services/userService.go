package services

import (
	"errors"
	"fmt"
	"time"
	"todo-user-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserDoesNotExists = errors.New("user does not exists")
)

type UserService struct {
	repo *repository.UsersRepo
}

func NewUserService(repo *repository.UsersRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(email string, password string) (int, string) {
	isUserExists := s.repo.CheckUserIsExists(email)
	if isUserExists {
		return 0, "user is already exists"
	}

	id, errorMessage := s.repo.CreateUser(email, password)
	return id, errorMessage
}

func (s *UserService) CreateSession(userId int, token string, expired_at time.Time) int {
	id := s.repo.CreateSession(userId, token, expired_at)
	return id
}

func (s *UserService) SignInUser(email string, password string) (string, error) {

	user, err := s.repo.GetByEmail(email)
	if err != nil || user.Email == "" {
		return "", ErrUserDoesNotExists
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	if err != nil {
		return "", ErrInvalidPassword
	}

	usr, err := s.repo.GetByEmail(email)
	fmt.Printf("%v", user)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(hashedPassword)
	fmt.Println(usr)
	return "asdas", nil
}
