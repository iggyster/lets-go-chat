package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iggyster/lets-go-chat/internal/user"
)

func TestNewRegister(t *testing.T) {
	var h Any = ProvideRegister(user.ProvideInMemoryUserRepo())
	_, ok := h.(*Register)
	if !ok {
		t.Errorf("faile to create new register handler")
	}
}

func TestRegister_ServerHTTP(t *testing.T) {
	var b bytes.Buffer
	var repo user.UserRepo = user.ProvideInMemoryUserRepo()

	json.NewEncoder(&b).Encode(AuthRequest{Username: "test", Password: "123qweasd"})

	req := httptest.NewRequest(http.MethodPost, "/user", &b)
	resp := httptest.NewRecorder()

	h := ProvideRegister(repo)
	h.ServeHTTP(resp, req)

	got := resp.Result().StatusCode
	want := http.StatusOK
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRegister_Validate(t *testing.T) {
	var repo user.UserRepo = user.ProvideInMemoryUserRepo()

	repo.Save(user.New("test", "secret-password"))

	data := []struct {
		name, json string
	}{
		{name: "Empty", json: `{"userName": "", "password": ""}`},
		{name: "Short", json: `{"userName": "te", "password": "passwor"}`},
		{name: "Exist", json: `{"userName": "test", "password": "secret-password"}`},
	}

	for _, input := range data {
		t.Run(input.name, func(t *testing.T) {
			jsonBody := []byte(input.json)

			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(jsonBody))
			resp := httptest.NewRecorder()

			h := ProvideRegister(repo)
			h.ServeHTTP(resp, req)

			got := resp.Result().StatusCode
			want := http.StatusBadRequest
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
