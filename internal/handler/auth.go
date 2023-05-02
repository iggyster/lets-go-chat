package handler

import (
	"encoding/json"
	"github.com/iggyster/lets-go-chat/internal/user"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
	"net/http"
	"time"
)

type LoginData struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type LoginResource struct {
	Url string `json:"url"`
}

func Auth(resp http.ResponseWriter, req *http.Request) {
	data := LoginData{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(resp, "Failed decoding", http.StatusInternalServerError)
		return
	}

	u := user.Repository.FindByUsername(data.Username)
	if !hasher.CheckPasswordHash(data.Password, u.Password) {
		http.Error(resp, "Invalid credentials", http.StatusBadRequest)
		return
	}

	resp.Header().Set("X-Rate-Limit", "5000")
	resp.Header().Set("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	err = json.NewEncoder(resp).Encode(LoginResource{Url: "ws://fancy-chat.io/ws&token=one-time-token"})
	if err != nil {
		http.Error(resp, "Failed encoding", http.StatusInternalServerError)
		return
	}
}
