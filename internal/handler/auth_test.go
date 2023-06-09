package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iggyster/lets-go-chat/internal/user"
)

type Any interface{}

func TestNewAuth(t *testing.T) {
	var h Any = NewAuth(user.NewRepo())

	_, ok := h.(*Auth)
	if !ok {
		t.Errorf("faile to create new auth handler")
	}
}

func TestAuth_ServerHTTP(t *testing.T) {
	var b bytes.Buffer
	var repo user.UserRepo = user.NewRepo()

	repo.Save(user.New("test", "123qweasd"))

	json.NewEncoder(&b).Encode(AuthRequest{Username: "test", Password: "123qweasd"})

	req, _ := http.NewRequest(http.MethodPost, "/user/login", &b)
	resp := httptest.NewRecorder()

	h := NewAuth(repo)
	h.ServeHTTP(resp, req)

	got := resp.Result().StatusCode
	want := http.StatusOK
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
