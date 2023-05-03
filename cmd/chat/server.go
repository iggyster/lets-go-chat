package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Accept)
	app.Use(logger.New())
	app.Use(recover.New())

	app.Post("/user", handler.Register)
	app.Post("/user/login", handler.Auth)

	log.Fatal(app.Listen(":8080"))
}
