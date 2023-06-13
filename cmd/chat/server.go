package main

import (
	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	"github.com/iggyster/lets-go-chat/internal/user"
)

var repo user.UserRepo = user.NewRepo()

func main() {
	app := app.New(":8080")

	app.Use(middleware.Recover)
	app.Use(middleware.Logger)

	app.Post("/user", handler.NewRegister(repo))
	app.Post("/user/login", handler.NewAuth(repo))
	app.Get("/ws", handler.NewChat(repo))
	app.Get("/users/active", handler.NewActive(repo))

	app.Start()
}
