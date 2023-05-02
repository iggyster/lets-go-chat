package user

import (
	"github.com/google/uuid"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
)

type User struct {
	Id, Username, Password string
}

func New(username, password string) *User {
	hashed, _ := hasher.HashPassword(password)

	return &User{
		Id:       uuid.New().String(),
		Username: username,
		Password: hashed,
	}
}
