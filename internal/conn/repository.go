package conn

import (
	"sync"

	"github.com/gorilla/websocket"
)

type (
	ConnRepo interface {
		Add(userId string, conn *websocket.Conn)
		Remove(userId string)
	}

	InMemoryConnRepo struct {
		sync.Map
	}
)

func NewConnRepo() ConnRepo {
	return &InMemoryConnRepo{}
}

func (repo *InMemoryConnRepo) Add(userId string, conn *websocket.Conn) {
	repo.Store(userId, conn)
}

func (repo *InMemoryConnRepo) Remove(userId string) {
	repo.Delete(userId)
}
