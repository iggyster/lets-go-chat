package handler

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/user"
)

func StartChat(conn *websocket.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close ws connection")
		}
	}()

	c := chat.New()

	token := conn.Query("token")
	if token == "" {
		log.Println("failed attemp to connect without token")
		return
	}

	usr := user.Repository.FindByToken(token)
	if usr == nil {
		log.Println("failed attemp to connect with invalid token")
		return
	}

	if !c.IsUserActivated(token) {
		c.AddUser(usr)
	}

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			c.RevokeToken(token)
			break
		}
		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			c.RevokeToken(token)
			break
		}
	}
}
