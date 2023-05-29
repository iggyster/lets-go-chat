package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	"github.com/iggyster/lets-go-chat/internal/router"
	"log"
)

func main() {
	app := fiber.New()

	middleware.Boot(app)
	router.Init(app)

	log.Fatal(app.Listen(":8080"))
}