package main

import (
	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/conn"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	"github.com/iggyster/lets-go-chat/internal/user"
)

var (
	userRepo user.UserRepo = user.NewRepo()
	connRepo conn.ConnRepo = conn.NewConnRepo()
)

func main() {
	app := app.New(":8080")

	app.Use(middleware.Recover)
	app.Use(middleware.Logger)

	app.Post("/user", handler.NewRegister(userRepo))
	app.Post("/user/login", handler.NewAuth(userRepo))
	app.Get("/ws", handler.NewChat(userRepo, connRepo))
	app.Get("/users/active", handler.NewActive(userRepo))

	app.Start()
}
