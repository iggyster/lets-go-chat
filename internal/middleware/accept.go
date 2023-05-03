package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Accept(ctx *fiber.Ctx) error {
	if ctx.Is("json") {
		return ctx.Next()
	}

	return fiber.NewError(fiber.StatusNotAcceptable, "Not acceptable content-type")
}
