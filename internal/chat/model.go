package chat

import (
	"sync"

	"github.com/iggyster/lets-go-chat/internal/user"
)

var (
	c    UserChat
	once sync.Once
)

type UserChat interface {
	AddUser(usr *user.User)
	GetUsers() []*user.User
	IsUserActivated(id string) bool
	DisconnectUser(usr *user.User)
}
type Chat struct {
	sync.Map
}

func New() UserChat {
	once.Do(func() {
		c = &Chat{}
	})

	return c
}

func (chat *Chat) AddUser(usr *user.User) {
	chat.Store(usr.Id, usr)
}

func (chat *Chat) IsUserActivated(id string) bool {
	_, ok := chat.Load(id)

	return ok
}

func (chat *Chat) GetUsers() []*user.User {
	var users []*user.User

	chat.Range(func(key, value any) bool {
		if usr, ok := value.(*user.User); ok {
			users = append(users, usr)
		}

		return true
	})

	return users
}

func (chat *Chat) DisconnectUser(usr *user.User) {
	chat.Delete(usr.Id)
}
