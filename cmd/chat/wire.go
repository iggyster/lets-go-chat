//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/iggyster/lets-go-chat/internal/app"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/db"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"github.com/iggyster/lets-go-chat/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func initializeDb() (*mongo.Client, func(), error) {
	wire.Build(db.ProvideClient)

	return nil, nil, nil
}

func initializeHub(db *mongo.Client) *chat.Hub {
	wire.Build(chat.MongoMessageRepoProviderSet, chat.NewHub)

	return &chat.Hub{}
}

func initializeHandlers(db *mongo.Client, hub *chat.Hub) *handler.HandlerProvider {
	wire.Build(chat.MongoMessageRepoProviderSet, user.InMemoryUserRepoProviderSet, handler.HandlerProviderSet)

	return &handler.HandlerProvider{}
}

func initializeApp(addr string) *app.App {
	wire.Build(app.ProvideApp)

	return &app.App{}
}
