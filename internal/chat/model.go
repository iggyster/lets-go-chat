package chat

import (
	"sync"

	"github.com/iggyster/lets-go-chat/internal/user"
)

var (
	c    IChat
	once sync.Once
)

type IChat interface {
	AddUser(usr *user.User)
	GetUsers() []*user.User
	IsUserActivated(token string) bool
	RevokeToken(token string)
}
type Chat struct {
	ActiveUsers sync.Map
}

func New() IChat {
	once.Do(func() {
		c = &Chat{ActiveUsers: sync.Map{}}
	})

	return c
}

func (chat *Chat) AddUser(usr *user.User) {
	chat.ActiveUsers.Store(usr.Token, usr)
}

func (chat *Chat) IsUserActivated(token string) bool {
	_, ok := chat.ActiveUsers.Load(token)

	return ok
}

func (chat *Chat) GetUsers() []*user.User {
	var users []*user.User

	chat.ActiveUsers.Range(func(key, value any) bool {
		if usr, ok := value.(*user.User); ok {
			users = append(users, usr)
		}

		return true
	})

	return users
}

func (chat *Chat) RevokeToken(token string) {
	chat.ActiveUsers.Delete(token)
}
