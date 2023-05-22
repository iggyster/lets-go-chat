package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/iggyster/lets-go-chat/internal/handler"
)

func Init(app *fiber.App) {
	app.Get("/ws", websocket.New(handler.StartChat))

	app.Get("/users/active", handler.GetUsers)

	app.Post("/user", handler.Register)
	app.Post("/user/login", handler.Auth)
}
