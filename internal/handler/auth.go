package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iggyster/lets-go-chat/internal/user"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
	"time"
)

type LoginData struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type LoginResource struct {
	Url string `json:"url"`
}

func Auth(ctx *fiber.Ctx) error {
	data := LoginData{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	usr := user.Repository.FindByUsername(data.Username)
	if usr != nil && !hasher.CheckPasswordHash(data.Password, usr.Password) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid credentials")
	}

	ctx.Append("X-Rate-Limit", "5000")
	ctx.Append("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())
	ctx.Append("Content-Type", "application/json")

	return ctx.Status(fiber.StatusOK).JSON(LoginResource{Url: "ws://fancy-chat.io/ws&token=one-time-token"})
}
