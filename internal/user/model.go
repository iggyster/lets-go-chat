package user

import (
	"github.com/google/uuid"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"-"`
}

func (usr *User) SetToken(token string) {
	usr.Token = token
}

func New(username, password string) *User {
	hashed, _ := hasher.HashPassword(password)

	return &User{
		Id:       uuid.New().String(),
		Username: username,
		Password: hashed,
	}
}
