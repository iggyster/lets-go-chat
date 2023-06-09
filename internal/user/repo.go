package user

import (
	"fmt"
	"sync"
)

type UserRepo interface {
	FindByUsername(username string) (*User, error)
	FindByToken(toekn string) *User
	FindActivated() []*User
	IsExists(username string) bool
	Save(user *User)
}

type InMemoryUserRepo struct {
	sync.Map
}

func NewRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{}
}

func (repo *InMemoryUserRepo) FindByUsername(username string) (*User, error) {
	val, ok := repo.Load(username)
	if !ok {
		return nil, fmt.Errorf("no user found")
	}

	usr, ok := val.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid user type %T", usr)
	}

	return usr, nil
}

func (repo *InMemoryUserRepo) FindByToken(token string) *User {
	var usr *User

	repo.Range(func(key, value any) bool {
		temp, ok := value.(*User)
		if ok && temp.Token == token {
			usr = temp
			return false
		}

		return true
	})

	return usr
}

func (repo *InMemoryUserRepo) FindActivated() []*User {
	var activated []*User

	repo.Range(func(key, value any) bool {
		if usr, ok := value.(*User); ok && usr.IsActivated() {
			activated = append(activated, usr)
		}

		return true
	})

	return activated
}

func (repo *InMemoryUserRepo) IsExists(username string) bool {
	_, ok := repo.Load(username)

	return ok
}

func (repo *InMemoryUserRepo) Save(usr *User) {
	repo.Store(usr.Username, usr)
}
