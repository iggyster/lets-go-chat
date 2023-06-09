package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iggyster/lets-go-chat/internal/user"
)

type Active struct {
	Repo user.UserRepo
}

func (a *Active) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := a.Repo.FindActivated()

	json.NewEncoder(w).Encode(users)
}

func NewActive(repo user.UserRepo) *Active {
	return &Active{Repo: repo}
}
