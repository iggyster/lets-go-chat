package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/iggyster/lets-go-chat/internal/user"
)

var ws websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Chat struct {
	Repo user.UserRepo
}

func NewChat(repo user.UserRepo) *Chat {
	return &Chat{Repo: repo}
}

func (c *Chat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := ws.Upgrade(w, req, nil)
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	token := req.URL.Query().Get("token")
	if token == "" {
		log.Println("invalid token accepted")

		return
	}

	usr := c.Repo.FindByToken(token)
	if usr == nil {
		log.Println("access denied")

		return
	}

	usr.RevokeToken()

	if !usr.IsActivated() {
		usr.Activate()
	}

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			usr.Deactivate()

			return
		}

		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			usr.Deactivate()

			return
		}
	}
}
