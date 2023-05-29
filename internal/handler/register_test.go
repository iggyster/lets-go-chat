package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iggyster/lets-go-chat/internal/user"

	"github.com/iggyster/lets-go-chat/internal/handler"

	"github.com/iggyster/lets-go-chat/internal/router"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	jsonBody := []byte(`{"userName": "test1", "password": "password123"}`)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json")

	app := fiber.New()
	router.Init(app)

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	usr := &handler.RegisterResource{}
	err := json.Unmarshal(body, usr)

	assert.Nil(t, err)
	assert.NotEmpty(t, usr.Id)
	assert.Equal(t, "test1", usr.UserName)
}

func TestRegister_Error(t *testing.T) {
	jsonBody := []byte(`"`)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json")

	app := fiber.New()
	router.Init(app)

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestRegister_validate(t *testing.T) {
	user.Repository.Save(user.New("test", "secret-password"))

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
			req.Header.Add("Content-Type", "application/json")

			app := fiber.New()
			router.Init(app)

			resp, _ := app.Test(req)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
	}
}
