package user

import (
	"fmt"
	"sync"
)

type Repo interface {
	FindByUsername(username string) (*User, error)
	FindByToken(toekn string) *User
	IsExists(username string) bool
	Save(user *User)
}

type InMemoryRepo struct {
	sync.Map
}

var Repository Repo = &InMemoryRepo{}

func (repo *InMemoryRepo) FindByUsername(username string) (*User, error) {
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

func (repo *InMemoryRepo) FindByToken(token string) *User {
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

func (repo *InMemoryRepo) IsExists(username string) bool {
	_, ok := repo.Load(username)

	return ok
}

func (repo *InMemoryRepo) Save(usr *User) {
	repo.Store(usr.Username, usr)
}
