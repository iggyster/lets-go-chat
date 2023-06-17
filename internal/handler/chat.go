package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/iggyster/lets-go-chat/internal/conn"
	"github.com/iggyster/lets-go-chat/internal/user"
)

var ws websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Chat struct {
	userRepo user.UserRepo
	connRepo conn.ConnRepo
}

func NewChat(userRepo user.UserRepo, connRepo conn.ConnRepo) *Chat {
	return &Chat{userRepo: userRepo, connRepo: connRepo}
}

func (c *Chat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn := c.newConn(w, req)
	defer conn.Close()

	token := req.URL.Query().Get("token")
	if token == "" {
		log.Println("invalid token accepted")

		return
	}

	usr := c.userRepo.FindByToken(token)
	if usr == nil {
		log.Println("access denied")

		return
	}

	c.activate(usr, conn)

	//TODO load all messages from DB

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			c.deactivate(usr)

			return
		}

		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			c.deactivate(usr)

			return
		}
	}
}

func (c *Chat) newConn(w http.ResponseWriter, req *http.Request) *websocket.Conn {
	conn, err := ws.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func (c *Chat) activate(usr *user.User, conn *websocket.Conn) {
	usr.RevokeToken()
	if usr.IsActivated() {
		return
	}

	usr.Activate()

	c.connRepo.Add(usr.Id, conn)
}

func (c *Chat) deactivate(usr *user.User) {
	usr.Deactivate()

	c.connRepo.Remove(usr.Id)
}
