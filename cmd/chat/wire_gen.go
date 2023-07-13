// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/db"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func initializeDb() (*mongo.Client, func(), error) {
	client, cleanup, err := db.ProvideClient()
	if err != nil {
		return nil, nil, err
	}
	return client, func() {
		cleanup()
	}, nil
}

func initializeHub(db2 *mongo.Client) *chat.Hub {
	mongoMessageRepo := chat.ProvideMessageRepo(db2)
	hub := chat.NewHub(mongoMessageRepo)
	return hub
}

func initializeHandlers(db2 *mongo.Client, hub *chat.Hub) *handler.HandlerProvider {
	inMemoryUserRepo := user.ProvideInMemoryUserRepo()
	active := handler.ProvideActive(inMemoryUserRepo)
	auth := handler.ProvideAuth(inMemoryUserRepo)
	register := handler.ProvideRegister(inMemoryUserRepo)
	mongoMessageRepo := chat.ProvideMessageRepo(db2)
	chatHandler := handler.ProvideChat(inMemoryUserRepo, mongoMessageRepo, hub)
	handlerProvider := handler.ProvideHandlers(active, auth, register, chatHandler)
	return handlerProvider
}

func initializeApp(addr string) *app.App {
	appApp := app.ProvideApp(addr)
	return appApp
}