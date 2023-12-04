package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Phone    string
	Password []byte
}
