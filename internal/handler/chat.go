package handler

import (
	"github.com/gofiber/websocket/v2"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/user"
	"log"
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

	usr.RevokeToken()

	if !c.IsUserActivated(usr.Id) {
		c.AddUser(usr)
	}

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			c.DisconnectUser(usr)
			break
		}
		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			c.DisconnectUser(usr)
			break
		}
	}
}
