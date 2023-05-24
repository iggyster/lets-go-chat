package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iggyster/lets-go-chat/internal/chat"
)

type Response struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func GetUsers(ctx *fiber.Ctx) error {
	c := chat.New()
	
	users := c.GetUsers()

	ctx.Status(fiber.StatusOK).JSON(users)

	return nil
}
