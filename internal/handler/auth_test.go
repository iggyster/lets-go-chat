package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iggyster/lets-go-chat/internal/user"
	"github.com/iggyster/lets-go-chat/mocks"
	"github.com/stretchr/testify/mock"
)

type Any interface{}

func TestNewAuth(t *testing.T) {
	var h Any = ProvideAuth(user.ProvideInMemoryUserRepo())

	_, ok := h.(*Auth)
	if !ok {
		t.Errorf("faile to create new auth handler")
	}
}

func TestAuth_ServerHTTP(t *testing.T) {
	var b bytes.Buffer
	var repo user.UserRepo = user.ProvideInMemoryUserRepo()

	repo.Save(user.New("test", "123qweasd"))

	json.NewEncoder(&b).Encode(AuthRequest{Username: "test", Password: "123qweasd"})

	req, _ := http.NewRequest(http.MethodPost, "/user/login", &b)
	resp := httptest.NewRecorder()

	h := ProvideAuth(repo)
	h.ServeHTTP(resp, req)

	got := resp.Result().StatusCode
	want := http.StatusOK
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkAuth_ServeHTTP(b *testing.B) {
	var buffer bytes.Buffer

	json.NewEncoder(&buffer).Encode(AuthRequest{Username: "test", Password: "123qweasd"})

	req := httptest.NewRequest("GET", "/user/login", &buffer)
	w := httptest.NewRecorder()

	repo := mocks.NewUserRepo(b)
	repo.On("FindByUsername", mock.Anything).Return(user.New("test", "123qweasd"), nil)
	repo.On("Save", mock.Anything).Return()

	handler := &Auth{Repo: repo}

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, req)
	}
}
