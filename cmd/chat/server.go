package main

import (
	"net/http/pprof"
	"os"

	_ "github.com/iggyster/lets-go-chat/docs"
	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Let's Go Chat
//	@version		1.0
//	@description	Chat application to practise Golang

// @host		localhost:8080
// @BasePath	/
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
	app.Get("/swagger/doc.json", httpSwagger.Handler())

	app.Get("/profiler/cpu", app.Handler(pprof.Profile))
	app.Get("/profiler/mem", pprof.Handler("heap"))
	app.Get("/profiler/alloc", pprof.Handler("allocs"))

	app.Start()
}
