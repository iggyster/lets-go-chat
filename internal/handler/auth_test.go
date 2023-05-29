package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iggyster/lets-go-chat/internal/handler"

	"github.com/iggyster/lets-go-chat/internal/router"

	"github.com/gofiber/fiber/v2"

	"github.com/iggyster/lets-go-chat/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	usr := user.New("test", "pass")
	user.Repository.Save(usr)

	jsonBody := []byte(`{"userName": "test", "password": "pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json")

	app := fiber.New()
	router.Init(app)

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	url := &handler.LoginResource{}
	err := json.Unmarshal(body, url)

	assert.Nil(t, err)
	assert.NotEmpty(t, url.Url)
}

func TestAuth_Validation(t *testing.T) {
	jsonBody := []byte(`{"userName": "test", "password": "secret"}`)
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json")

	app := fiber.New()
	router.Init(app)

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAuth_Error(t *testing.T) {
	jsonBody := []byte(`"`)
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json")

	app := fiber.New()
	router.Init(app)

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
