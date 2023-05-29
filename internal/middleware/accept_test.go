package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAccept(t *testing.T) {
	app := fiber.New()
	app.Post("/", Accept, func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Ok!"})
	})

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)

	req = httptest.NewRequest(http.MethodPost, "/", nil)
	resp, err = app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotAcceptable, resp.StatusCode)
}
