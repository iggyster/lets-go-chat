package handler

import (
	"github.com/google/wire"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/user"
)

var HandlerProviderSet = wire.NewSet(
	ProvideActive,
	ProvideAuth,
	ProvideChat,
	ProvideRegister,
	ProvideHandlers,
)

type HandlerProvider struct {
	Active   *Active
	Auth     *Auth
	Register *Register
	Chat     *ChatHandler
}

func ProvideHandlers(
	active *Active,
	auth *Auth,
	register *Register,
	chat *ChatHandler,
) *HandlerProvider {
	return &HandlerProvider{
		Active:   active,
		Auth:     auth,
		Register: register,
		Chat:     chat,
	}
}

func ProvideActive(repo user.UserRepo) *Active {
	return &Active{Repo: repo}
}

func ProvideAuth(repo user.UserRepo) *Auth {
	return &Auth{Repo: repo}
}

func ProvideChat(
	userRepo user.UserRepo,
	messageRepo chat.MessageRepo,
	hub *chat.Hub,
) *ChatHandler {
	return &ChatHandler{
		userRepo:    userRepo,
		messageRepo: messageRepo,
		hub:         hub,
	}
}

func ProvideRegister(repo user.UserRepo) *Register {
	return &Register{Repo: repo}
}
