package handler

import (
	"log"
	"net/http"

	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/user"
)

type ChatHandler struct {
	userRepo    user.UserRepo
	messageRepo chat.MessageRepo
	hub         *chat.Hub
}

//	@Summary		Chat
//	@Description	Websocket to connect Chat application
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			token	query	string	true	"Access token to login to the chat"
//	@Success		101
//	@Router			/ws [get]
func (chatHandler *ChatHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")
	if token == "" {
		log.Println("invalid token accepted")

		return
	}

	usr := chatHandler.userRepo.FindByToken(token)
	if usr == nil {
		log.Println("access denied")

		return
	}

	conn := chat.NewConn(w, req)
	client := chat.NewClient(chatHandler.hub, conn, usr)

	messages := chatHandler.messageRepo.FindAll()
	if messages != nil {
		client.RestoreHistory(&messages)
	}

	go client.Read()
	go client.Write()
}
