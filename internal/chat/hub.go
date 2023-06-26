package chat

import (
	"sync"
)

type Hub struct {
	connect    chan *Client
	disconnect chan *Client
	broadcast  chan *Msg
	clients    sync.Map
	repo       MessageRepo
}

func NewHub(repo MessageRepo) *Hub {
	return &Hub{
		connect:    make(chan *Client),
		disconnect: make(chan *Client),
		broadcast:  make(chan *Msg),
		clients:    sync.Map{},
		repo:       repo,
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.connect:
			hub.clients.Store(client, true)
		case client := <-hub.disconnect:
			if _, ok := hub.clients.Load(client); ok {
				hub.disconnectClient(client)
			}
		case msg := <-hub.broadcast:
			hub.repo.Save(msg)
			hub.clients.Range(func(key, value any) bool {
				client, ok := key.(*Client)
				if !ok && msg.SentBy(client) {
					return true
				}

				client.send <- msg

				return true
			})
		}
	}
}

func (hub *Hub) disconnectClient(client *Client) {
	close(client.send)
	hub.clients.Delete(client)
}
