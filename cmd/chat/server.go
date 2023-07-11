package main

import (
	"os"

	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/middleware"
)

func main() {
	app.InitEnv()

	db, disconnect, _ := initializeDb()
	defer disconnect()

	hub := initializeHub(db)
	go hub.Run()

	handlers := initializeHandlers(db, hub)
	app := initializeApp(":" + os.Getenv("APP_PORT"))

	app.Use(middleware.Recover)
	app.Use(middleware.Logger)

	app.Post("/user", handlers.Register)
	app.Post("/user/login", handlers.Auth)
	app.Get("/ws", handlers.Chat)
	app.Get("/users/active", handlers.Active)

	app.Start()
}
