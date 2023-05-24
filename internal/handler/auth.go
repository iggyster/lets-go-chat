package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/iggyster/lets-go-chat/internal/user"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
	"github.com/iggyster/lets-go-chat/pkg/tokengenerator"
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

	usr, err := user.Repository.FindByUsername(data.Username)
	if err != nil && !hasher.CheckPasswordHash(data.Password, usr.Password) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid credentials")
	}

	token, err := tokengenerator.Generate(16)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Token generation failed")
	}

	usr.SetToken(token)
	user.Repository.Save(usr)

	ctx.Append("X-Rate-Limit", "5000")
	ctx.Append("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())

	return ctx.Status(fiber.StatusOK).JSON(LoginResource{Url: fmt.Sprintf("ws://localhost:8080/ws?token=%v", token)})
}
