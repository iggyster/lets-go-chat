package chat

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/iggyster/lets-go-chat/internal/user"
)

const (
	maxMessageSize = 512
)

var (
	ws = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	usr  *user.User
	send chan *Msg
}

func NewClient(hub *Hub, conn *websocket.Conn, usr *user.User) *Client {
	client := &Client{
		hub:  hub,
		conn: conn,
		usr:  usr,
		send: make(chan *Msg),
	}

	hub.connect <- client

	usr.Activate()

	return client
}

func NewConn(w http.ResponseWriter, req *http.Request) *websocket.Conn {
	conn, err := ws.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func (client *Client) RestoreHistory(messages *[]Msg) {
	for _, msg := range *messages {
		client.conn.WriteMessage(websocket.TextMessage, msg.GetMessage())
	}
}

func (client *Client) Read() {
	defer client.disconnect()

	client.conn.SetReadLimit(maxMessageSize)

	for {
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("read error: %v\n", err)
			}
			break
		}

		prefix := []byte(fmt.Sprintf("%v: ", client.usr.Username))
		msg = append(prefix, bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))...)

		client.hub.broadcast <- NewMessage(client, string(msg))

		log.Printf("received: %v\n", string(msg))
	}
}

func (client *Client) Write() {
	defer client.disconnect()

	for {
		select {
		case msg, ok := <-client.send:
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})

				return
			}

			if msg.SentBy(client) {
				continue
			}

			err := client.conn.WriteMessage(websocket.TextMessage, msg.GetMessage())
			if err != nil {
				log.Printf("write error: %v\n", err)

				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.conn.Close()
	client.usr.Deactivate()

	client.hub.disconnect <- client
}
