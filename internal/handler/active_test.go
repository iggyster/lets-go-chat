package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iggyster/lets-go-chat/internal/user"
)

func TestNewActive(t *testing.T) {
	var h Any = NewActive(user.NewRepo())
	_, ok := h.(*Active)
	if !ok {
		t.Errorf("faile to create new active handler")
	}
}

func TestActive_ServerHTTP(t *testing.T) {
	var repo user.UserRepo = user.NewRepo()

	usr := user.New("test", "123qweasd")
	usr.Activate()
	repo.Save(usr)

	req, _ := http.NewRequest(http.MethodGet, "/users/active", nil)
	resp := httptest.NewRecorder()

	h := NewActive(repo)
	h.ServeHTTP(resp, req)

	got := resp.Result().StatusCode
	want := http.StatusOK
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
