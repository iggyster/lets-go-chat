package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/iggyster/lets-go-chat/internal/user"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
	"github.com/iggyster/lets-go-chat/pkg/tokengen"
)

type (
	Auth struct {
		Repo user.UserRepo
	}

	AuthRequest struct {
		Username string `json:"userName"`
		Password string `json:"password"`
	}

	AuthResponse struct {
		Url string `json:"url"`
	}
)

func (handler *Auth) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := decodeAuthRequest(w, req)

	usr, err := handler.Repo.FindByUsername(data.Username)
	if err != nil || !hasher.CheckPasswordHash(data.Password, usr.Password) {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}

	token := tokengen.New(16)
	usr.SetToken(token)
	handler.Repo.Save(usr)

	w.Header().Set("X-Rate-Limit", "5000")
	w.Header().Set("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())

	host := os.Getenv("APP_HOST")

	json.NewEncoder(w).Encode(AuthResponse{Url: fmt.Sprintf("ws://%v:8080/ws?token=%v", host, token)})
}

func decodeAuthRequest(w http.ResponseWriter, req *http.Request) AuthRequest {
	var data AuthRequest

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return data
}
