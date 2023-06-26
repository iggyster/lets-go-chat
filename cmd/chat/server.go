package main

import (
	"context"

	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/db"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	"github.com/iggyster/lets-go-chat/internal/user"
)

var (
	userRepo    user.UserRepo
	messageRepo chat.MessageRepo
)

func main() {
	db := db.NewMongoClient()
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	userRepo = user.NewRepo()
	messageRepo = chat.NewRepo(db)

	hub := chat.NewHub(messageRepo)
	app := app.New(":8080")

	go hub.Run()

	app.Use(middleware.Recover)
	app.Use(middleware.Logger)

	app.Post("/user", handler.NewRegister(userRepo))
	app.Post("/user/login", handler.NewAuth(userRepo))
	app.Get("/ws", handler.NewChatHandler(userRepo, messageRepo, hub))
	app.Get("/users/active", handler.NewActive(userRepo))

	app.Start()
}
