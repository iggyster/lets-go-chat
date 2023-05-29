package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoot(t *testing.T) {
	app := fiber.New()
	Boot(app)

	assert.Equal(t, uint32(3), app.HandlersCount())
}
